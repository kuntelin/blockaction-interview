package auth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")
	if header == "" {
		return "", errors.New("no 'Authorization' header")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	logger.Debug("ExtractBearerToken: " + jwtToken[1])

	return jwtToken[1], nil
}
