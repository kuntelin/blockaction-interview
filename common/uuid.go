package common

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	var uuidStr string
	uuidInstant, uuidErr := uuid.NewV7()
	if uuidErr != nil {
		uuidStr = uuidInstant.String()
	} else {
		uuidStr = uuid.New().String()
	}
	return uuidStr
}
