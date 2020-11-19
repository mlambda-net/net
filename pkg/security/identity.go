package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
)

type Identity interface {
	Authenticate() bool

}

type identity struct {
	token *jwt.Token
}

func (i identity) Authenticate() bool {
	claims, ok := i.token.Claims.(jwt.MapClaims)
	if ok && i.token.Valid && claims["authorize"].(bool) {
		return true
	}
	return false

}

func NewIdentity(text string) (Identity, error)  {
	token, err := getToken(getBearer(text))
	if err != nil {
		return nil, err
	}
	return &identity{token: token}, nil
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