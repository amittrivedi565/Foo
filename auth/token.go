package token

import (
	"Foo/types"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("foobar")


func IssueToken(userid string, role string) (string, error){
	claims := types.Token{
		Id: userid,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}