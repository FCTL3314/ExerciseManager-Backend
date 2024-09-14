package domain

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type JwtCustomRefreshClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
