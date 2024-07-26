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

func GetSetting() *Setting {
	return &Setting{
		DEBUG:        strings.EqualFold(os.Getenv("DEBUG"), "true"),
		PORT:         os.Getenv("PORT"),
		LOG_LEVEL:    os.Getenv("LOG_LEVEL"),
		LOG_OUTPUT:   os.Getenv("LOG_OUTPUT"),
		DATABASE_URL: os.Getenv("DATABASE_URL"),
	}
}
