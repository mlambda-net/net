package security

type Roles struct {
  App string `json:"app"`
  Name string `json:"name"`
}

type User struct {
  Audience string `json:"aud"`
  Email string `json:"email"`
  Expiration string `json:"exp"`
  Issuer string `json:"iss"`
  Name string `json:"name"`
  Subject string `json:"sub"`
  Roles []Roles `json:"roles"`
}
