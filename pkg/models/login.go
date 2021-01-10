package models

type Login struct {
	UserID int    `json:"id"`
	Token  string `json:"token"`
}
