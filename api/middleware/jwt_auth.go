package middleware

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsErrorResponse)
			c.Abort()
			return
		}

		token := tokenParts[1]
		schema := tokenParts[0]
		authorized, _ := tokenutil.IsAuthorized(token, secret)

		if !authorized || schema != "Bearer" {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsErrorResponse)
			c.Abort()
			return
		}

		userID, err := tokenutil.ExtractIDFromToken(token, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsErrorResponse)
			c.Abort()
			return
		}
		c.Set("x-user-id", userID)
		c.Next()
	}
}
