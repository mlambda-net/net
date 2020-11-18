package security

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Cors(t *testing.T)  {

	rq, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	handler := Cors(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(rr, rq)

	assert.Equal(t, "*", rr.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "POST, GET, OPTIONS, PUT, DELETE", rr.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization", rr.Header().Get("Access-Control-Allow-Headers"))

}