package common

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"

	_ "github.com/joho/godotenv/autoload"
	"github.com/op/go-logging"
)

func GetLogger() *logging.Logger {
	setting := GetSetting()

	level := setting.LOG_LEVEL
	output := setting.LOG_OUTPUT

	logger := logging.MustGetLogger("blockaction-api")

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

func initLogger2(_ string, level string, output string) *zap.Logger {
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

func logging_example() {
	logger1 := initLogger1("blockaction-api", os.Getenv("LOG_LEVEL"), os.Getenv("LOG_OUTPUT"))
	logger1.Debug("debug")
	logger1.Info("info")
	logger1.Notice("notice")
	logger1.Warning("warning")
	logger1.Error("err")
	logger1.Critical("crit")

	logger2 := initLogger2("blockaction-api", os.Getenv("LOG_LEVEL"), os.Getenv("LOG_OUTPUT"))
	logger2.Debug("debug")
	logger2.Info("info")
	logger2.Warn("warn")
	logger2.Error("err")
	logger2.DPanic("dpanic")
	logger2.Panic("panic")
	logger2.Fatal("fatal")
}
