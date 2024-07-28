package auth

import (
	"github.com/google/uuid"
)

func uuidGenerator() string {
	uuidValue, _ := uuid.NewV7()

	return uuidValue.String()
}

func CreateTokenService(username *string) Token {
	token := Token{
		Token:    *username + "::" + uuidGenerator(),
		Username: *username,
	}

	createTokenResult := db.Model(&Token{}).Create(&token)
	if createTokenResult.Error != nil {
		return Token{}
	}

	return token
}

func GetTokenService(tokenValue *string) (Token, error) {
	var token Token

	getTokenResult := db.Where("token = ?", *tokenValue).Find(&token)
	if getTokenResult.Error != nil {
		return Token{}, getTokenResult.Error
	}

	return token, nil
}

func DeleteTokenService(tokenValue *string) {
	db.Delete(&Token{}, tokenValue)
}
