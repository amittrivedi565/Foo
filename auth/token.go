package auth

import (
    "Foo/types"
    "fmt"

    "github.com/golang-jwt/jwt/v5"
)

var secret = []byte("foobar")

// Issue a JWT with custom claims
func IssueToken(userid string, role string) (string, error) {
    claims := types.Token{
        Id:   userid,
        Role: role,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secret)
}

// Validate and return claims payload
func ValidateToken(tokenString string) (*types.Token, error) {
    token, err := jwt.ParseWithClaims(tokenString, &types.Token{}, func(token *jwt.Token) (interface{}, error) {
        return secret, nil
    })

    if err != nil {
        return nil, fmt.Errorf("parse error: %w", err)
    }

    if claims, ok := token.Claims.(*types.Token); ok && token.Valid {
        return claims, nil
    }
    
    return nil, fmt.Errorf("invalid token")
}
