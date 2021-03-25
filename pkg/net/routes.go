package net

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/security"
  "net/http"
)

type Route interface {
  AddRoute(name string, path string, isSecure bool, method string, handler func(w http.ResponseWriter, r *http.Request))
  Add(func(router * mux.Router))
  GetRouter() *mux.Router
}

type route struct {
  name string
  path string
  isSecure bool
  method string
  handler func(w http.ResponseWriter, r *http.Request)
}

type router struct {
  routes []route
  config *metrics.Configuration
  extend func(router *mux.Router)
}

func (r *router) Add(router func(router *mux.Router)) {
  r.extend = router
}

func (r *router) GetRouter() *mux.Router {
  router := mux.NewRouter()
  m := metrics.NewMetric(r.config)
  router.Use(m.Trace)
  for _, v := range r.routes {
    m.AddMetric(fmt.Sprintf("%s/%s", v.path, v.method), v.name)
    if v.isSecure {
      router.Handle(v.path,security.Authenticate(http.HandlerFunc(v.handler))).Methods(v.method)
    } else {
      router.Handle(v.path,http.HandlerFunc(v.handler)).Methods(v.method)
    }
  }
  if r.extend != nil {
    r.extend(router)
  }

  return router
}

func (r *router) AddRoute(name string, path string, isSecure bool, method string, handler func(w http.ResponseWriter, r *http.Request)) {
  r.routes = append(r.routes, route{
    name:    name,
    path:     path,
    isSecure: isSecure,
    method: method,
    handler: handler,
  })
}

func NewRoute(config *metrics.Configuration ) Route {
  return &router{ config : config }
}
