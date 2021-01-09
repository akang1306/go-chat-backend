package controller

import (
	"github.com/challenge/pkg/models"
	"github.com/challenge/pkg/service"
)

type serviceHandler interface {
	CreateUser(models.User) (models.User, error)
	Health() models.Health
}

// Handler provides the interface to handle different requests
type Handler struct {
	Service serviceHandler
}

func NewController() Handler {
	return Handler{Service: service.NewService()}
}
