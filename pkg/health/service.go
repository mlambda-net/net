package health

import (
  "fmt"
  "github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
  "github.com/mlambda-net/net/pkg/metrics"
)

type Service interface {
	Start(r *mux.Router) func(opts ...healthcheck.Option)
  SetConfig(config *metrics.Configuration)
}

type service struct {
  port   int32
  config *metrics.Configuration
}

func (s *service) SetConfig(config *metrics.Configuration) {
 s.config = config
}

func(s *service) Start( r *mux.Router) func (opts ...healthcheck.Option) {
  return func(opts ...healthcheck.Option) {
    r.Handle(fmt.Sprintf("%s/healthz", s.config.App.Path), s.healthy(opts...)).Name("healthz")
    r.HandleFunc(fmt.Sprintf("%s/live", s.config.App.Path), s.live()).Name("live")
  }
}

func NewHealthServer (port int32) Service  {
	return &service{port: port}
}
