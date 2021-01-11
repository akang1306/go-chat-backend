package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/challenge/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

var defaultExpiration time.Duration = time.Hour
var secretKey string = "secret"

type Manager struct {
}

// TokenForUser returns a token valid for the given user with default expiration.
func (Manager) TokenForUser(user *models.User) string {
	expiresAt := time.Now().Add(defaultExpiration).Unix()

	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["exp"] = expiresAt

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

// VerifyToken checks the tokenString and returns userID if valid, error if not.
func (Manager) VerifyToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}
	userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(userID), nil
}

func New() Manager {
	return Manager{}
}
