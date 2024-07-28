package users

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
	user := User{
		Username: username,
		Password: EncryptPassword(password),
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
