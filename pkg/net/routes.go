package net

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/mlambda-net/net/pkg/metrics"
  "github.com/mlambda-net/net/pkg/security"
  "github.com/sirupsen/logrus"
  "github.com/throttled/throttled/v2"
  "github.com/throttled/throttled/v2/store/memstore"
  "net/http"
)

type Route interface {
  AddRoute(name string, path string, query []string, isSecure bool, method string, handler func(w http.ResponseWriter, r *http.Request))
  Add(func(router * mux.Router))
  GetRouter() *mux.Router
}

type route struct {
  name string
  path string
  isSecure bool
  method string
  query []string
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

  store, err := memstore.New(65536)
  if err != nil {
    logrus.Fatal(err)
  }

  quota := throttled.RateQuota{
    MaxRate:  throttled.PerMin(20),
    MaxBurst: 5,
  }
  rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
  if err != nil {
    logrus.Fatal(err)
  }

  httpRateLimiter := throttled.HTTPRateLimiter{
    RateLimiter: rateLimiter,
    VaryBy:      &throttled.VaryBy{Path: true},
  }

  for _, v := range r.routes {
    m.AddMetric(fmt.Sprintf("%s/%s", v.path, v.method), v.name)
    if v.isSecure {
      if v.query != nil {
        router.Handle(v.path,httpRateLimiter.RateLimit( security.Authenticate(http.HandlerFunc(v.handler)))).Methods(v.method)
      } else {
        router.Handle(v.path, httpRateLimiter.RateLimit(security.Authenticate(http.HandlerFunc(v.handler)))).Methods(v.method).Queries(r.queryPars(v)...)
      }
    } else {
      if v.query != nil {
        router.Handle(v.path, httpRateLimiter.RateLimit(http.HandlerFunc(v.handler))).Methods(v.method)
      }else {
        router.Handle(v.path, httpRateLimiter.RateLimit(http.HandlerFunc(v.handler))).Methods(v.method).Queries(r.queryPars(v)...)
      }
    }
  }
  if r.extend != nil {
    r.extend(router)
  }

  return router
}

func (r *router) queryPars(v route) []string {
  qs := make([]string, 0)
  for _, q := range v.query {
    qs = append(qs, q)
    qs = append(qs, fmt.Sprintf("{%s}", q))
  }
  return qs
}

func (r *router) AddRoute(name string, path string, query []string, isSecure bool, method string, handler func(w http.ResponseWriter, r *http.Request)) {
  r.routes = append(r.routes, route{
    name:    name,
    path:     path,
    isSecure: isSecure,
    method: method,
    handler: handler,
    query: query,
  })
}

func NewRoute(config *metrics.Configuration ) Route {
  return &router{ config : config }
}
