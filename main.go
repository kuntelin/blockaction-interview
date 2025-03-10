package main

import (
	"blockaction-api/app/auth"
	"blockaction-api/app/users"
	"blockaction-api/common"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/op/go-logging"
	"gorm.io/gorm"
)

var setting = common.GetSetting()
var logger *logging.Logger = common.GetLogger()
var db *gorm.DB = common.GetDB()

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
	server.GET("/health", func(c *gin.Context) {
		var uuid string

		// check database connection
		healthyResult := db.Raw("SELECT gen_random_uuid();")
		if healthyResult.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": healthyResult.Error.Error(),
			})
			c.Abort()
			return
		}

		// check database function
		healthyResult.Scan(&uuid)
		if uuid == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to get uuid",
			})
			c.Abort()
			return
		}

		// connection and database function is healthy
		c.JSON(http.StatusOK, gin.H{
			"message": "ok, uuid: " + uuid,
		})
		c.Abort()
	})

	auth_r := server.Group("")
	{
		logger.Debug("Registering auth routes")
		// register auth routes
		auth.Init(db)
		auth.RegisterRoutes(auth_r, "/auth")
	}

	api_v1_r := server.Group("/api/v1")
	api_v1_r.Use(auth.AuthMiddleware())
	{
		logger.Debug("Registering user routes")
		// register user routes
		users.Init(db)
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
