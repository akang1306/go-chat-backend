package service

import (
	"github.com/challenge/pkg/models"
)

// CreateUser attempts to create user and returns user if successful, returns error if not.
func (h Handler) CreateUser(user models.User) (models.User, error) {
	return user, nil
}
