package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	setting := GetSetting()

	db, err := gorm.Open(postgres.Open(setting.DATABASE_URL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
