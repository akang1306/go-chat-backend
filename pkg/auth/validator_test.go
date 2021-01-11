package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/challenge/pkg/auth"
)

type mockValidator struct {
	IsValid bool
}

func (v mockValidator) VerifyToken(tokenString string) (int, error) {
	if !v.IsValid {
		return 0, errors.New("error")
	}
	return 0, nil
}

type handler struct{}

func TestUserValidator(t *testing.T) {
	invalidTokenHandler := auth.NewValidateUserHandler(mockValidator{IsValid: false})
	validTokenHandler := auth.NewValidateUserHandler(mockValidator{IsValid: true})

	var handlerFnCalled bool = false

	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		handlerFnCalled = true
	}
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("", "/test", nil)

	// Invalid token should avoid handler call.
	invalidTokenHandler(handlerFn).ServeHTTP(writer, request)
	if handlerFnCalled {
		t.Errorf("Error: expected handler to not be called but was")
	}

	// Valid token should call handler.
	validTokenHandler(handlerFn).ServeHTTP(writer, request)
	if !handlerFnCalled {
		t.Errorf("Error: expected handler to be called but was not")
	}
}
