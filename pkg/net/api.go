package net

import (
	"fmt"
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
	"github.com/mlambda-net/net/pkg/health"
	"github.com/mlambda-net/net/pkg/metrics"
	"github.com/mlambda-net/net/pkg/security"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type Api interface {
  Register(f func(s Route))
	Metrics(f func(c *metrics.Configuration))
	Checks(options ...healthcheck.Option)
	Start()
	Wait()
}

type api struct {
  port   int32
  health int32
  config *metrics.Configuration

  options []healthcheck.Option
  sem     chan int
  maps   map[string]string
  routes Route
}



func (a *api) Register(f func(s Route)) {
	a.routes = NewRoute(a.config)
	f(a.routes)
}

func (a *api) Wait() {
	<-a.sem
}

func (a *api) Start() {

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
		routes := a.routes.GetRouter()
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.port), security.Cors(routes)))
		a.sem <- 1
	}()

}

func (a *api) Metrics(f func(c *metrics.Configuration)) {
	a.config = &metrics.Configuration{}
	f(a.config)
}

func (a *api) Checks(options ...healthcheck.Option) {
	a.options = options
}

func NewApi(port int32, health int32) Api {
	return &api{
		port: port,
		health: health,
		maps: make(map[string]string),

	}
}
