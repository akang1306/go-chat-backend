package models

type Login struct {
	UserID UserID `json:"id"`
	Token  string `json:"token"`
}
