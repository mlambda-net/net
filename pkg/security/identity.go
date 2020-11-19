package security

import (
	"os"
	"strings"
)

type Identity interface {
	Authenticate() bool

}

type identity struct {
	claims Claims
}

func (i identity) Authenticate() bool {
	return i.claims.Get("authorize").(bool)
}

func NewIdentity(text string) (Identity, error)  {
	token := NewToken(os.Getenv("SECRET_KEY"))
	claims, err := token.Claims(getBearer(text))
	if err != nil {
		return nil, err
	}

	return &identity{claims: claims}, nil
}


func getBearer( bearer string ) string  {
	items := strings.Split(bearer, " ")

	if len(items) == 2 {

		auth := items[1]
		return auth
	}
	return ""
}
