package pkg

import (
	"fmt"
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
	"github.com/mlambda-net/net/pkg/health"
	"github.com/mlambda-net/net/pkg/metrics"
	"github.com/mlambda-net/net/pkg/security"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

type Api interface {
	Metrics(f func(c metrics.Configuration))
	Checks(options ...healthcheck.Option)
	RegisterAuth(f func(r *mux.Router))
	RegisterWithAuth(f func(r *mux.Router))
	Start()
	Wait()
	Trace(route string, name string)
}

type api struct {
	port    int32
	health  int32
	config  metrics.Configuration
	secure  *mux.Router
	routes  *mux.Router
	options []healthcheck.Option
	sem     chan int
	maps    map[string]string
}

func (a api) Trace(route string, name string) {
	a.maps[route] = name
}

func (a api) Wait() {
	<-a.sem
}

func (a api) Start() {

	a.sem = make(chan int, 2)

	go func() {
		r := mux.NewRouter()
		s := health.NewHealthServer(a.health)
		s.Start(r)(a.options...)
		r.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.health), r))
		a.sem <- 1
	}()

	go func() {
		m := metrics.NewMetric(a.config, a.maps)
		an := negroni.New(negroni.Wrap(m.Trace(security.Authenticate(a.secure))), negroni.Wrap(m.Trace(a.routes)))
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), an))
		a.sem <- 1
	}()

}

func (a api) RegisterAuth(f func(r *mux.Router)) {
	f(a.secure)
}

func (a api) RegisterWithAuth(f func(r *mux.Router)) {
	f(a.routes)
}

func (a api) Metrics(f func(c metrics.Configuration)) {
	a.config = metrics.Configuration{}
	f(a.config)
}

func (a api) Checks(options ...healthcheck.Option) {
	a.options = options
}

func NewApi(port int32, health int32) Api  {
	return api{
		port: port,
		health: health,
		secure : mux.NewRouter().StrictSlash(true),
		routes : mux.NewRouter().StrictSlash(true),
		maps: make(map[string]string),
	}
}
