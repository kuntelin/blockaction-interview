package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, prefix string) {
	s := r.Group(prefix)
	{
		s.GET("", ListUserController)
		s.POST("", CreateUserController)
		s.GET("/:username", GetUserController)
	}
}
