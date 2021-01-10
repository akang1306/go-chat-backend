package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/challenge/pkg/jwt"
)

type tokenValidator interface {
	VerifyToken(tokenString string) (int, error)
}

const authorizationHeader string = "Authorization"

// ValidateUser checks for a token and validates it
// before allowing the method to execute
func ValidateUser(handler http.HandlerFunc) http.HandlerFunc {
	validator := jwt.New()
	return func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(r.Header.Get("Authorization"))
		id, err := validator.VerifyToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", id)))
	}
}
