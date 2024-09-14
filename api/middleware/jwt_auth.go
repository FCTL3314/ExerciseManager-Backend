package middleware

import (
	"ExerciseManager/internal/domain"
	"ExerciseManager/internal/tokenutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
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
		authorized, _ := tokenutil.IsAuthorized(token, secret)

		if !authorized || schema != "Bearer" {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
			c.Abort()
			return
		}

		userIDString, err := tokenutil.ExtractIDFromToken(token, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
			c.Abort()
			return
		}

		userID, err := strconv.ParseUint(userIDString, 10, 64)
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.InvalidAuthCredentialsResponse)
		}

		c.Set("x-user-id", uint(userID))
		c.Next()
	}
}
