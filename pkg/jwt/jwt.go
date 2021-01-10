package jwt

import (
	"time"

	"github.com/challenge/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

var defaultExpiration time.Duration = time.Hour
var secretKey string = "secret"

type Manager struct {
}

type Token struct {
	UserID int
	*jwt.StandardClaims
}

// TokenForUser returns a token valid for the given user with default expiration.
func (Manager) TokenForUser(user *models.User) string {
	expiresAt := time.Now().Add(defaultExpiration).Unix()
	tokenClaims := &Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenClaims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

// VerifyToken checks the tokenString and returns userID if valid, error if not.
func (Manager) VerifyToken(tokenString string) (int, error) {
	token := Token{}
	_, err := jwt.ParseWithClaims(tokenString, token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	return token.UserID, nil
}

func New() Manager {
	return Manager{}
}