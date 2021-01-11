package controller

import (
	"net/http"

	"github.com/challenge/pkg/helpers"
)

// Check returns the health of the service and DB
func (h Handler) Check(w http.ResponseWriter, r *http.Request) {
	helpers.RespondJSON(w, h.Service.Health(r.Context()))
}
