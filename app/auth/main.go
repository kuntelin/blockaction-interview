package auth

import (
	"blockaction-api/common"

	"gorm.io/gorm"
)

var logger = common.GetLogger()
var db *gorm.DB

func Init(d *gorm.DB) {
	logger.Debug("Initializing auth module")

	db = d
	RegisterModels()
}

func RegisterModels() error {
	db.AutoMigrate(&Token{})
	return nil
}
