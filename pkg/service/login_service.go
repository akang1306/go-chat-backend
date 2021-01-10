package service

import (
	"errors"

	"github.com/challenge/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const invalidLoginMsg string = "invalid login credentials"

// Login checks if user exists and returns a valid login model if true, returns error if not.
func (h Handler) Login(username, password string) (models.Login, error) {
	user, err := h.db.GetUserByUsername(username)
	if err != nil {
		return models.Login{}, errors.New(invalidLoginMsg)
	}
	if errCompare := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); errCompare != nil {
		return models.Login{}, errors.New(invalidLoginMsg)
	}

	return models.Login{
		UserID: user.ID,
		Token:  h.tokenManager.TokenForUser(user),
	}, nil
}
