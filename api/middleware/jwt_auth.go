package middleware

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/tokenutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(tokenManager tokenutil.IJWTTokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
			c.Abort()
			return
		}

		token := tokenParts[1]
		schema := tokenParts[0]
		authorized, _ := tokenManager.IsAccessTokenValid(token)

		if !authorized || schema != "Bearer" {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
			c.Abort()
			return
		}

		userId, err := tokenManager.ExtractUserIDFromAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
			c.Abort()
			return
		}

		c.Set("x-user-id", userId)
		c.Next()
	}
}
