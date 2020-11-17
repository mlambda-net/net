package health

import (
	"context"
	"github.com/etherlabsio/healthcheck"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_Healthy(t *testing.T) {

	s := NewHealthServer(8080)
	r := mux.NewRouter()
	s.Start(r)(
		healthcheck.WithTimeout(5*time.Second),
		healthcheck.WithChecker("google", healthcheck.CheckerFunc(func(ctx context.Context) error {
			_, err := http.Get("http://google.com")
			return err
		})),
	)

	req, _ := http.NewRequest("GET", "/healthz", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)


}
