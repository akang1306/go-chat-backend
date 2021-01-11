package service

import (
	"context"
	"errors"

	"github.com/challenge/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

const invalidLoginMsg string = "invalid login credentials"

// Login checks if user exists and returns a valid login model if true, returns error if not.
func (h Handler) Login(ctx context.Context, username, password string) (*models.Login, error) {
	user, err := h.db.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, errors.New(invalidLoginMsg)
	}
	if errCompare := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); errCompare != nil {
		return nil, errors.New(invalidLoginMsg)
	}

	return &models.Login{
		UserID: user.ID,
		Token:  h.tokenManager.TokenForUser(user),
	}, nil
}
