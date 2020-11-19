package metrics

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_RegisterMetrics(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/user", nil)
	assert.Nil(t, err)

	cf := &Configuration{}
	cf.App.Name = "sample"
	m := metric{
		config: cf,
	}

	rr := httptest.NewRecorder()
	handler := m.Trace(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(rr, req)
	body :=  rr.Body.String()
	println(body)

}
