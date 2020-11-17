package health

import (
	"github.com/etherlabsio/healthcheck"
	"net/http"
)

func (s *service) healthy( opts ...healthcheck.Option ) http.Handler {
	return healthcheck.Handler(opts...)
}



