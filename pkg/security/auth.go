package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

func Authenticate( next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token, err := getToken(getBearer(r.Header.Get("Authorization")))

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				authorized := claims["authorized"].(bool)
				if authorized {
					next.ServeHTTP(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		}

	})
}

func getBearer( bearer string ) string  {
	items := strings.Split(bearer, " ")

	if len(items) == 2 {

		auth := items[1]
		return auth
	}
	return ""
}

func getToken(text string)  (*jwt.Token, error ) {
	token, err := jwt.Parse(text, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}