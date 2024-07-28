package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddlewareImpl(c *gin.Context) {
	logger.Debug("Calling AuthMiddlewareImpl")

	// check authorizaiton header is given
	headerToken, extractErr := ExtractBearerToken(c)
	if extractErr != nil {
		logger.Error("extractErr: " + extractErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
		})
		c.Abort()
		return
	}

	// check token exists
	tokenInstance, getTokenErr := GetTokenService(&headerToken)
	if getTokenErr != nil {
		logger.Error("getTokenErr: " + getTokenErr.Error())
		c.JSON(http.StatusForbidden, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
		})
		c.Abort()
		return
	}

	if tokenInstance.Token == "" {
		logger.Error("tokenInstance.Token is empty")
		c.JSON(http.StatusForbidden, gin.H{
			"return_code": 1001,
			"msgid":       "unauthorized",
			"msgdata":     nil,
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
