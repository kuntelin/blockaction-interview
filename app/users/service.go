package users

import (
	"errors"
)

var users = []User{
	{
		ID:       1,
		Username: "admin",
		Password: "admin",
		Email:    "admin@example.com",
		IsAdmin:  true,
	},
	{
		ID:       2,
		Username: "alex",
		Password: "alex",
		Email:    "alex@example.com",
		IsAdmin:  false,
	},
}

func checkUserExists(username string) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}

func ListUserService() []User {
	return users
}

func CreateUserService(username string, password string, email string) (User, error) {
	if checkUserExists(username) {
		return User{}, errors.New("User already exists")
	}

	user := User{
		ID:       len(users) + 1,
		Username: username,
		Password: password,
		Email:    email,
		IsAdmin:  false,
	}
	users = append(users, user)
	return user, nil
}

func GetUserService(username string) User {
	for _, u := range users {
		if u.Username == username {
			return u
		}
	}
	return User{}
}
