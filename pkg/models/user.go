package models

type UserID = int64

type User struct {
	ID       UserID `json:"id"`
	Username string `json:"-"`
	Password []byte `json:"-"`
}
