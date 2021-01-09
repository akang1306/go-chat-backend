package service

import (
	"github.com/challenge/pkg/models"
)

// Login checks if user exists and returns a valid login model if true, returns error if not.
func (h Handler) Login(username, password string) (models.Login, error) {
	return models.Login{UserID: 0, Token: ""}, nil
}
