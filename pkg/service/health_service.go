package service

import "github.com/challenge/pkg/models"

func (h Handler) Health() models.Health {
	return models.Health{Health: "ok"}
}
