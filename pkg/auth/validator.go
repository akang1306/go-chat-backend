package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/challenge/pkg/models"
)

type tokenValidator interface {
	VerifyToken(tokenString string) (models.UserID, error)
}

const authorizationHeader string = "Authorization"

// NewValidateUserHandler returns a function that checks for a token and validates it
// before allowing the method to execute
func NewValidateUserHandler(validator tokenValidator) func(http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := strings.TrimSpace(r.Header.Get(authorizationHeader))
			id, err := validator.VerifyToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", id)))
		}
	}
}
