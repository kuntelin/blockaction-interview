package users

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strconv"
)

func EncryptPassword(password string) string {
	// sha256 password
	h := sha256.New()
	h.Write([]byte(password))

	// salted sha256 password
	salt := 1639992878
	_, _ = io.WriteString(h, strconv.Itoa(salt))

	sum := h.Sum(nil)
	s := hex.EncodeToString(sum)

	return string(s)
}

func ValidatePassword(plainPassword *string, encryptedPassword *string) bool {
	return EncryptPassword(*plainPassword) == *encryptedPassword
}
