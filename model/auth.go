package model

type AuthToken struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}
