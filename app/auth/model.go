package auth

type Token struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
