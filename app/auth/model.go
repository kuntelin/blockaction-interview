package auth

import (
	"time"
)

// type Token struct {
// 	Token    string `json:"token"`
// 	Username string `json:"username"`
// }

type Token struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null;DEFAULT:current_timestamp;"`
	Token     string    `json:"token" gorm:"primaryKey"`
	Username  string    `json:"username"`
}
