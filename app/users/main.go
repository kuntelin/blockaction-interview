package users

import (
	"blockaction-api/common"

	"gorm.io/gorm"
)

var logger = common.GetLogger()
var db *gorm.DB

func Init(d *gorm.DB) {
	logger.Debug("Initializing users module")
	db = d

	RegisterModels()
	Bootstrap()
}

func RegisterModels() error {
	db.AutoMigrate(&User{})
	return nil
}

func Bootstrap() {
	logger.Debug("Adding default users")

	var defaultUsers = []User{
		{
			Username: "admin",
			Password: "admin",
			Email:    "admin@example.com",
			IsAdmin:  true,
		},
		{
			Username: "user",
			Password: "user",
			Email:    "user@example.com",
			IsAdmin:  false,
		},
	}

	for _, u := range defaultUsers {
		_, createErr := CreateUserService(u.Username, u.Password, u.Email)
		if createErr != nil {
			logger.Warningf("Failed to insert default user: %v\n", createErr)
		} else {
			logger.Warningf(
				"Default user inserted: %s, default password is '%s', please change it after signin\n",
				u.Username,
				u.Password,
			)
		}
	}
}
