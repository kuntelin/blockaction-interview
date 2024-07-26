package common

import (
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type Setting struct {
	DEBUG        bool
	PORT         string
	LOG_LEVEL    string
	LOG_OUTPUT   string
	DATABASE_URL string
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func GetSetting() *Setting {
	return &Setting{
		DEBUG:        strings.EqualFold(getEnv("DEBUG", "false"), "true"),
		PORT:         getEnv("PORT", "8080"),
		LOG_LEVEL:    getEnv("LOG_LEVEL", "WARNING"),
		LOG_OUTPUT:   getEnv("LOG_OUTPUT", "ext://sys.stdout"),
		DATABASE_URL: getEnv("DATABASE_URL", ""),
	}
}
