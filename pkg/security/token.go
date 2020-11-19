package security

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token interface {
	Create(map[string]interface{}) (string,error)
	Claims(result string) (Claims,error)
}

type Claims interface {
	Get(name string) interface{}
	ToMap() map[string]string
}

type claims struct {
	values jwt.MapClaims
}

func (c claims) ToMap() map[string]string {
	maps := make(map[string]string)
	for k, v := range c.values {
		maps[k] = fmt.Sprintf("%v", v)
	}
	return maps
}

func (c claims) Get(name string) interface{} {
	return c.values[name]
}

type token struct {
	secret string
}



func (t token) Create(values map[string]interface{}) (string, error) {
	claims :=  jwt.MapClaims{}
	for k, v:= range values {
		claims[k]= v
	}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (t token) Claims(text string) (Claims,error) {
	token, err := t.getToken(text)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("the token is invalid")
	}
	return &claims{values: token.Claims.(jwt.MapClaims)}, nil
}

func NewToken(secret string) Token  {
	return &token{secret: secret}
}

func (t token) getToken(text string)  (*jwt.Token, error ) {
	token, err := jwt.Parse(text, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.secret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}