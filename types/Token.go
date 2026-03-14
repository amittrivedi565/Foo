package types

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	Id   string
	Role string
	jwt.RegisteredClaims

}

