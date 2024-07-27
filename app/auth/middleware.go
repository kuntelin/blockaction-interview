package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddlewareImpl(c *gin.Context) {
	logger.Debug("Message from AuthMiddlewareImpl")

	// check authorizaiton header is given
	headerToken, extractErr := ExtractBearerToken(c.GetHeader("Authorization"))
	if extractErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
			"trace":       extractErr.Error(),
		})
		c.Abort()
		return
	}
	logger.Debug("headerToken is " + headerToken)

	// check token exists
	tokenInstance, getTokenErr := GetTokenService(&headerToken)
	if getTokenErr != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
			"trace":       getTokenErr.Error(),
		})
		c.Abort()
		return
	}

	logger.Debug("tokenInstance is " + tokenInstance.Token)

	c.Next()
}

func AuthMiddleware() gin.HandlerFunc {
	return authMiddlewareImpl
}
