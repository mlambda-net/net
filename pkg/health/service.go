package health

import (
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
)

type Service interface {
	Start(r *mux.Router) func(opts ...healthcheck.Option)
}

type service struct {
	port int32
}

func(s *service) Start( r *mux.Router) func (opts ...healthcheck.Option) {
	return func(opts ...healthcheck.Option) {
		r.Handle("/healthz", s.healthy(opts...)).Name("healthz")
		r.HandleFunc("/live", s.live()).Name("live")
	}
}

func NewHealthServer (port int32) Service  {
	return &service{port: port}
}
