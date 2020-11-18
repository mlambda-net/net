package pkg

import (
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
  "github.com/mlambda-net/net/pkg/common"
  "github.com/mlambda-net/net/pkg/metrics"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_ApiLoad(t *testing.T){

	api := NewApi(8080, 9090)
	api.Metrics(func( c metrics.Configuration) {
		c.App.Name = "app"
		c.App.Env = "dev"
		c.App.Version = "1.0.0"
		c.Metric.Namespace = "ns"
		c.Metric.SubSystem = "ss"
	})



	api.Register(func(r common.Route) {
	 r.AddRoute("a", "/api/a",  true, func(w http.ResponseWriter, _ *http.Request) {
     w.WriteHeader(200)
   } )

	 r.AddRoute("b", "/api/b", false, func(w http.ResponseWriter, _ *http.Request) {
     w.WriteHeader(200)
   } )

  })

	api.Checks( healthcheck.WithTimeout(5*time.Second) )
	api.Start()

 	r, e := http.Get("http://localhost:8080/api/b")
	assert.Nil(t, e)
 	assert.Equal(t, "200 OK", r.Status)

	r, e = http.Get("http://localhost:8080/api/a")
	assert.Nil(t, e)
	assert.Equal(t, "401 Unauthorized", r.Status)

	r, e = http.Get("http://localhost:9090/metrics")
	body, _ := ioutil.ReadAll(r.Body)
	println(string(body))
	assert.Nil(t, e)
	assert.Equal(t, "200 OK", r.Status)


}
