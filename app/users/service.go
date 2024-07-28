package users

var usersRespository = []User{
	// {
	// 	ID:       common.GenerateUUID(),
	// 	Username: "admin",
	// 	Password: "admin",
	// 	Email:    "admin@example.com",
	// 	IsAdmin:  true,
	// },
	// {
	// 	ID:       common.GenerateUUID(),
	// 	Username: "alex",
	// 	Password: "alex",
	// 	Email:    "alex@example.com",
	// 	IsAdmin:  false,
	// },
}

func checkUserExists(username *string) bool {
	result := false

	logger.Debugf("Check User Exists: %s\n", *username)
	dbResult := db.Find(&User{}).Where("username = ?", username)
	if dbResult.Error != nil {
		logger.Errorf("Check User Exists Failed: %v\n", dbResult.Error)
	} else {
		result = true
	}
	return result
}

func ListUserService() []User {
	var users []User

	logger.Debug("Get User List")
	dbResult := db.Find(&users)
	if dbResult.Error != nil {
		logger.Errorf("Get User List Failed: %v\n", dbResult.Error)
	}
	return users
}

func CreateUserService(username string, password string, email string) (User, error) {
	// if checkUserExists(&username) {
	// 	return User{}, errors.New("User already exists")
	// }

	user := User{
		// ID:       common.GenerateUUID(),
		Username: username,
		Password: password,
		Email:    email,
		IsAdmin:  false,
	}
	// usersRespository = append(usersRespository, user)
	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("Create User Failed: %v\n", err)
		return User{}, err
	}

	return user, nil
}

func GetUserService(username *string) User {
	var user User

	logger.Debugf("Get User Info: %s\n", *username)
	dbResult := db.Where("username = ?", username).Find(&user)
	if dbResult.Error != nil {
		logger.Errorf("Get User Info Failed: %v\n", dbResult.Error)
	}
	return user
}
