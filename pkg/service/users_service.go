package service

import (
	"context"
	"errors"

	"github.com/challenge/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser attempts to create user with given username & password and returns user if successful, returns error if not.
func (h Handler) CreateUser(ctx context.Context, username, password string) (models.User, error) {
	user := models.User{}
	if username == "" || password == "" {
		return user, errors.New("empty username or password")
	}
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		return user, hashErr
	}
	user.Username = username
	user.Password = hash
	id, err := h.db.AddUser(ctx, user)
	user.ID = id
	return user, err
}
