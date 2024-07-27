package users

import (
	"blockaction-api/common"
	"errors"
)

var usersRespository = []User{
	{
		ID:       common.GenerateUUID(),
		Username: "admin",
		Password: "admin",
		Email:    "admin@example.com",
		IsAdmin:  true,
	},
	{
		ID:       common.GenerateUUID(),
		Username: "alex",
		Password: "alex",
		Email:    "alex@example.com",
		IsAdmin:  false,
	},
}

func checkUserExists(username *string) bool {
	for _, u := range usersRespository {
		if u.Username == *username {
			return true
		}
	}
	return false
}

func ListUserService() []User {
	return usersRespository
}

func CreateUserService(username string, password string, email string) (User, error) {
	if checkUserExists(&username) {
		return User{}, errors.New("User already exists")
	}

	user := User{
		ID:       common.GenerateUUID(),
		Username: username,
		Password: password,
		Email:    email,
		IsAdmin:  false,
	}
	usersRespository = append(usersRespository, user)
	return user, nil
}

func GetUserService(username *string) User {
	for _, u := range usersRespository {
		if u.Username == *username {
			return u
		}
	}
	return User{}
}
