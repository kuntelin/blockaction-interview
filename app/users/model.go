package users

import (
	"time"
)

// type User struct {
// 	ID       string `json:"id" gorm:"primary_key"`
// 	Username string `json:"username" gorm:"unique"`
// 	Password string `json:"-"`
// 	Email    string `json:"email"`
// 	IsAdmin  bool   `json:"is_admin"`
// }

type User struct {
	// gorm.Model
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime;not null;DEFAULT:current_timestamp;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime;not null;DEFAULT:current_timestamp;"`
	ID        string    `json:"id" gorm:"type:uuid;DEFAULT:gen_random_uuid();primary_key;"`
	Username  string    `json:"username" gorm:"size:100;unique;index;not null;"`
	Password  string    `json:"-" gorm:"size:100;not null;"`
	Email     string    `json:"email" gorm:"size:100;index;not null;"`
	IsAdmin   bool      `json:"is_admin" gorm:"DEFAULT:false;"`
}
