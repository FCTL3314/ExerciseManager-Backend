package tokenutil

import (
	"ExerciseManager/internal/domain"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type TokenManager struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

func NewTokenManager(accessSecret, refreshSecret string, accessExpiry, refreshExpiry time.Duration) *TokenManager {
	return &TokenManager{
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
		AccessExpiry:  accessExpiry,
		RefreshExpiry: refreshExpiry,
	}
}

func (tm *TokenManager) CreateAccessToken(user *domain.User) (string, error) {
	exp := time.Now().Add(tm.AccessExpiry)
	claims := &domain.JwtCustomAccessClaims{
		ID: strconv.FormatUint(uint64(user.ID), 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return tm.createToken(claims, tm.AccessSecret)
}

func (tm *TokenManager) CreateRefreshToken(user *domain.User) (string, error) {
	exp := time.Now().Add(tm.RefreshExpiry)
	claims := &domain.JwtCustomRefreshClaims{
		ID: strconv.FormatUint(uint64(user.ID), 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	return tm.createToken(claims, tm.RefreshSecret)
}

func (tm *TokenManager) IsAccessTokenValid(tokenStr string) (bool, error) {
	claims := &domain.JwtCustomAccessClaims{}
	token, err := tm.parseToken(tokenStr, tm.AccessSecret, claims)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (tm *TokenManager) IsRefreshTokenValid(tokenStr string) (bool, error) {
	claims := &domain.JwtCustomRefreshClaims{}
	token, err := tm.parseToken(tokenStr, tm.RefreshSecret, claims)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (tm *TokenManager) ExtractUserIDFromAccessToken(tokenStr string) (string, error) {
	claims := &domain.JwtCustomAccessClaims{}
	token, err := tm.parseToken(tokenStr, tm.AccessSecret, claims)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}

func (tm *TokenManager) ExtractUserIDFromRefreshToken(tokenStr string) (string, error) {
	claims := &domain.JwtCustomRefreshClaims{}
	token, err := tm.parseToken(tokenStr, tm.RefreshSecret, claims)
	if err != nil {
		return "", err
	}

	if !token.Valid && !errors.Is(err, jwt.ErrTokenExpired) {
		return "", fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}

func (tm *TokenManager) createToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (tm *TokenManager) parseToken(tokenStr, secret string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}
