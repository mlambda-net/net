package common

import (
  "github.com/gorilla/mux"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/security"
  "net/http"
)

type Route interface {
  AddRoute(name string, path string, isSecure bool, handler func(w http.ResponseWriter, r *http.Request))
  GetRouter() *mux.Router
}

type route struct {
  name string
  path string
  isSecure bool
  handler func(w http.ResponseWriter, r *http.Request)
}

type router struct {
  routes []route
  config metrics.Configuration
}

func (r *router) GetRouter() *mux.Router {
  router := mux.NewRouter()
  m := metrics.NewMetric(r.config)

  for _, v := range r.routes {
    m.AddMetric(v.path, v.name)
    if v.isSecure {
      router.Handle(v.path, m.Trace(security.Authenticate(http.HandlerFunc(v.handler))))
    } else {
      router.Handle(v.path, m.Trace(http.HandlerFunc(v.handler)))
    }
  }
  return router
}




func (r *router) AddRoute(name string, path string, isSecure bool, handler func(w http.ResponseWriter, r *http.Request)) {
  r.routes = append(r.routes, route{
    name:    name,
    path:     path,
    isSecure: isSecure,
    handler: handler,
  })
}

func NewRoute(config metrics.Configuration ) Route  {
  return &router{ config : config }
}
