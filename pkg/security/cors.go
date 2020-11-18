package security

import "net/http"

func Cors( next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupHeaders(&w, r)
		if (*r).Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setupHeaders(w *http.ResponseWriter, _ *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}