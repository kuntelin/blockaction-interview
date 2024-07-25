package users

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}
