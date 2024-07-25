package main

import (
	"blockaction-api/app/users"
	"blockaction-api/common"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/op/go-logging"
)

var setting = common.GetSetting()
var logger *logging.Logger = common.GetLogger()

func main() {
	if setting.DEBUG {
		logger.Info("Runing in debug mode")
		gin.SetMode(gin.DebugMode)
	} else {
		logger.Debug("Runing in release mode")
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()

	// container health check
	server.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "health check",
		})
	})

	api_v1 := server.Group("/api/v1")
	{
		logger.Debug("Registering user routes")
		// register user routes
		users.Init()
		users.RouteUsers(api_v1, "/users")
	}

	// blog := &routes.Blog{}
	// server.GET("/blogs", blog.GetBlogs)
	// server.GET("/blogs/:id", blog.GetBlog)
	// server.POST("/blogs", blog.CreateBlog)
	// server.PUT("/blogs/:id", blog.UpdateBlog)
	// server.DELETE("/blogs/:id", blog.DeleteBlog)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}
