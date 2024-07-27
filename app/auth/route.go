package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, prefix string) {
	s := r.Group(prefix)
	{
		s.POST("/signin", SignInUserController)
		s.POST("/signout", SignOutUserController)
		s.POST("/signup", SignUpUserController)
	}
}
