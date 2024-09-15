package tokenutil

import (
	"ExerciseManager/internal/domain"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type IJWTTokenManager interface {
	CreateUserAccessToken(user *domain.User) (string, error)
	CreateUserRefreshToken(user *domain.User) (string, error)
	IsAccessTokenValid(tokenStr string) (bool, error)
	IsRefreshTokenValid(tokenStr string) (bool, error)
	ExtractUserIDFromAccessToken(tokenStr string) (int64, error)
	ExtractUserIDFromRefreshToken(tokenStr string) (int64, error)
}

type JWTTokenManager struct {
	AccessSecret  string
	RefreshSecret string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

func NewJWTTokenManager(accessSecret, refreshSecret string, accessExpiry, refreshExpiry time.Duration) *JWTTokenManager {
	return &JWTTokenManager{
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
		AccessExpiry:  accessExpiry,
		RefreshExpiry: refreshExpiry,
	}
}

func (tm *JWTTokenManager) CreateUserAccessToken(user *domain.User) (string, error) {
	return tm.CreateAccessToken(user.ID)
}

func (tm *JWTTokenManager) CreateUserRefreshToken(user *domain.User) (string, error) {
	return tm.CreateRefreshToken(user.ID)
}

func (tm *JWTTokenManager) IsAccessTokenValid(tokenStr string) (bool, error) {
	return tm.isTokenValid(tokenStr, tm.AccessSecret)
}

func (tm *JWTTokenManager) IsRefreshTokenValid(tokenStr string) (bool, error) {
	return tm.isTokenValid(tokenStr, tm.RefreshSecret)
}

func (tm *JWTTokenManager) isTokenValid(tokenStr string, secret string) (bool, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := tm.parseToken(tokenStr, secret, claims)
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (tm *JWTTokenManager) ExtractUserIDFromAccessToken(tokenStr string) (int64, error) {
	return tm.ExtractUserIDFromToken(tokenStr, tm.AccessSecret)
}

func (tm *JWTTokenManager) ExtractUserIDFromRefreshToken(tokenStr string) (int64, error) {
	return tm.ExtractUserIDFromToken(tokenStr, tm.RefreshSecret)
}

func (tm *JWTTokenManager) ExtractUserIDFromToken(tokenStr string, secret string) (int64, error) {
	idString, err := tm.extractIdFromToken(tokenStr, secret)
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return 0, domain.ErrInvalidAuthCredentials
	}

	return id, nil
}

func (tm *JWTTokenManager) extractIdFromToken(tokenStr string, secret string) (string, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := tm.parseToken(tokenStr, secret, claims)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims.ID, nil
}

func (tm *JWTTokenManager) CreateAccessToken(id int64) (string, error) {
	return tm.createToken(id, tm.AccessSecret)
}

func (tm *JWTTokenManager) CreateRefreshToken(id int64) (string, error) {
	return tm.createToken(id, tm.RefreshSecret)
}

func (tm *JWTTokenManager) createToken(id int64, secret string) (string, error) {
	exp := time.Now().Add(tm.AccessExpiry)
	claims := &jwt.RegisteredClaims{
		ID:        strconv.FormatInt(id, 10),
		ExpiresAt: jwt.NewNumericDate(exp),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (tm *JWTTokenManager) parseToken(tokenStr, secret string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return token, err
}
