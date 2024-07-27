package auth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func authMiddlewareImpl(c *gin.Context) {
	logger.Debug("Message from AuthMiddlewareImpl")
	jwtToken, extractErr := extractBearerToken(c.GetHeader("Authorization"))
	if extractErr != nil {
		c.JSON(401, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
			"trace":       extractErr.Error(),
		})
		c.Abort()
		return
	}

	logger.Debug("jwtToken", jwtToken)
	c.Next()
}

func AuthMiddleware() gin.HandlerFunc {
	return authMiddlewareImpl
}
