package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var tokenRepository = []Token{}

func CreateTokenService(username *string) Token {
	var tokenValue string
	uuidValue, err := uuid.NewV7()
	if err == nil {
		tokenValue = uuidValue.String()
	} else {
		tokenValue = time.Now().Format(time.RFC3339)
	}
	token := Token{
		Token:    *username + "::" + tokenValue,
		Username: *username,
	}
	tokenRepository = append(tokenRepository, token)
	return token
}

func GetTokenService(tokenValue *string) (Token, error) {
	for _, t := range tokenRepository {
		if t.Token == *tokenValue {
			return t, nil
		}
	}
	return Token{}, errors.New("token not found")
}

func DeleteTokenService(tokenValue *string) error {
	var newTokens []Token
	for _, t := range tokenRepository {
		if t.Token != *tokenValue {
			newTokens = append(newTokens, t)
		}
	}
	tokenRepository = newTokens
	return nil
}
