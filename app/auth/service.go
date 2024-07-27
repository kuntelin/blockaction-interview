package auth

import (
	"errors"
	"time"
)

var tokens = []Token{}

func checkTokenExist(token *string) bool {
	for _, t := range tokens {
		if t.Token == *token {
			return true
		}
	}
	return false
}

func CreateTokenService(username *string) Token {
	token := Token{
		Token:    *username + "::" + time.Now().Format(time.RFC3339),
		Username: *username,
	}
	tokens = append(tokens, token)
	return token
}

func GetTokenService(tokenValue *string) (Token, error) {
	for _, t := range tokens {
		if t.Token == *tokenValue {
			return t, nil
		}
	}
	return Token{}, errors.New("token not found")
}

func DeleteTokenService(tokenValue *string) error {
	var newTokens = []Token{}
	for _, t := range tokens {
		if t.Token != *tokenValue {
			newTokens = append(newTokens, t)
		}
	}
	tokens = newTokens
	return nil
}
