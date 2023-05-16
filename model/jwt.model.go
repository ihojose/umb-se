package model

import "gopkg.in/square/go-jose.v2/jwt"

type Token struct {
	Token string `json:"token"`
}

type Payload struct {
	Username uint   `json:"username"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Role     int32  `json:"role"`
}

type CustomClaims struct {
	*jwt.Claims
	PrivateClaim1      string                 `json:"privateClaim1,omitempty"`
	PrivateClaim2      []string               `json:"privateClaim2,omitEmpty"`
	AnyJSONObjectClaim map[string]interface{} `json:"anyJSONObjectClaim"`
}
