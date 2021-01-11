package service

import (
	"context"
	"errors"

	"github.com/challenge/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser attempts to create user with given username & password and returns user if successful, returns error if not.
func (h Handler) CreateUser(ctx context.Context, username, password string) (models.UserID, error) {
	if username == "" || password == "" {
		return 0, errors.New("empty username or password")
	}
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return 0, hashErr
	}
	return h.db.AddUser(ctx, username, hash)
}
