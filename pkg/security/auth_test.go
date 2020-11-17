package security

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Auth( t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU2MzgwMjAsInVzZXJfaWQiOiJjb3lvdGVAYWNtZS5jb20ifQ.wLQX8pmzqqxQU5RWxRo6hbQmLfQeiKQs9iFhx5p3Czg")

	rr := httptest.NewRecorder()
	handler := Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(rr,req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_Fail(t *testing.T)  {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "bearer eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU2MzgwMjAsInVzZXJfaWQiOiJjb3lvdGVAYWNtZS5jb20ifQ")

	rr := httptest.NewRecorder()
	handler := Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))

	handler.ServeHTTP(rr,req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
