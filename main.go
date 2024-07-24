package main

import (
	"fmt"
	"os"

	"blockaction-api/routes"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/op/go-logging"
	"go.uber.org/zap"
)

var dsn = os.Getenv("DATABASE_URL")

func initLogger1(name string, level string, output string) *logging.Logger {
	logger := logging.MustGetLogger(name)

	formatter := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{module}.%{shortfunc} ▶ [%{level:-8s}] %{color:reset}%{message}`,
	)

	// get the log level from the environment variable or default to ERROR
	log_level, err := logging.LogLevel(level)
	if err != nil {
		log_level = logging.ERROR
	}

	// create a new backend using the default writer
	var logger_io *os.File
	if output == "" || strings.EqualFold(output, "ext://sys.stdout") {
		logger_io = os.Stdout
	} else {
		logger_io, err = os.OpenFile(output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Printf("Failed to create '%s' file, using console output\n", output)
			logger_io = os.Stdout
		}
	}
	backend := logging.NewLogBackend(logger_io, "", 0)

	// set the backends format to the formatter
	backendFormatter := logging.NewBackendFormatter(backend, formatter)

	// set the backend to only log messages with the level
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(log_level, "")

	// Set the backends to be used.
	logging.SetBackend(backendLeveled)

	return logger
}

func initLogger2(name string, level string, output string) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zap_level, err := zap.ParseAtomicLevel(level)
	if err != nil {
		zap_level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	}
	config.Level = zap_level
	config.OutputPaths = []string{output}

	logger, err := config.Build()
	if err != nil {
		logger, _ = zap.NewDevelopment()
	}

	return logger
}

func main() {
	logger := initLogger1("blockaction-api", os.Getenv("LOG_LEVEL"), os.Getenv("LOG_OUTPUT"))
	logger.Debug("debug")
	logger.Info("info")
	logger.Notice("notice")
	logger.Warning("warning")
	logger.Error("err")
	logger.Critical("crit")

	// logger := initLogger2("blockaction-api", os.Getenv("LOG_LEVEL"), os.Getenv("LOG_OUTPUT"))
	// logger.Debug("debug")
	// logger.Info("info")
	// logger.Warn("warn")
	// logger.Error("err")
	// logger.DPanic("dpanic")
	// logger.Panic("panic")
	// logger.Fatal("fatal")

	if dsn == "" {
		logger.Error("DATABASE_URL is not set")
		os.Exit(1)
	}
	logger.Debug(fmt.Sprintf("DATABASE_URL: %s", dsn))

	router := gin.Default()

	blog := &routes.Blog{}
	router.GET("/blogs", blog.GetBlogs)
	router.GET("/blogs/:id", blog.GetBlog)
	router.POST("/blogs", blog.CreateBlog)
	router.PUT("/blogs/:id", blog.UpdateBlog)
	router.DELETE("/blogs/:id", blog.DeleteBlog)

	router.Run(":8080")
}
