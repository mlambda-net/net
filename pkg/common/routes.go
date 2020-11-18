package common

import "net/http"

type Route interface {
  AddRoute(name string, path string, isSecure bool, handler func(w http.ResponseWriter, r *http.Request))

}

type route struct {

}

func (r route) AddRoute(name string, path string, isSecure bool, handler func(w http.ResponseWriter, r *http.Request)) {

}

func NewRoutes() Route  {
  return &route{}
}
