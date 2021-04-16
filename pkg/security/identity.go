package security

import (
  "fmt"
  "os"
  "strings"
)

type Identity interface {
	Authenticate() bool
	GetHeaders() string
  HasRoles(roles []string) bool
}

type identity struct {
	claims Claims
}

func (i *identity) HasRoles(roles []string) bool {

  if roles == nil {
    return true
  }

  c := i.claims.Get("roles")
  if c != nil {
    r := c.([]interface{})
    for _, rs := range r {
      k := rs.(map[string]interface{})
      role := fmt.Sprintf("%s-%s", k["app"], k["name"])

      for _, ro := range roles {
        if strings.ToLower(role) == strings.ToLower(ro) {
          return true
        }
      }
    }
  }

  return false

}

func (i *identity) GetHeaders() string {
	return i.claims.ToString()
}

func (i *identity) Authenticate() bool {
	return i.claims.Get("sub") != ""
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

