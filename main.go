package main

import (
	"blockaction-api/app/auth"
	"blockaction-api/app/users"
	"blockaction-api/common"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/op/go-logging"
)

var setting = common.GetSetting()
var logger *logging.Logger = common.GetLogger()

func RequestPrintHelloworld() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Helloworld")
		c.Next()
	}
}

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
			"message": "ok",
		})
	})

	auth_r := server.Group("")
	{
		logger.Debug("Registering auth routes")
		// register auth routes
		auth.Init()
		auth.RegisterRoutes(auth_r, "/auth")
	}

	api_v1_r := server.Group("/api/v1")
	api_v1_r.Use(auth.AuthMiddleware())
	{
		logger.Debug("Registering user routes")
		// register user routes
		users.Init()
		users.RegisterRoutes(api_v1_r, "/users")
	}

	// blog := &routes.Blog{}
	// server.GET("/blogs", blog.GetBlogs)
	// server.GET("/blogs/:id", blog.GetBlog)
	// server.POST("/blogs", blog.CreateBlog)
	// server.PUT("/blogs/:id", blog.UpdateBlog)
	// server.DELETE("/blogs/:id", blog.DeleteBlog)

	// * server without gracefull shutdown
	// err := server.Run(":" + setting.PORT)
	// if err != nil {
	// 	panic(err)
	// }

	// * server with gracefull shutdown
	srv := &http.Server{
		Addr:    ":" + setting.PORT,
		Handler: server.Handler(),
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Warning("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Warning("Server Shutdown:", err)
	}
	logger.Warning("Server exiting")
}
