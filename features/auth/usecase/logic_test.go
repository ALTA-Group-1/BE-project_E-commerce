package usecase

import (
	"errors"
	"project/e-commerce/features/auth"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginAuthorized(t *testing.T) {
	repo := new(mocks.AuthData)
	// dataInput := auth.Core{Email: "fakhry@mail.id", Password: "fakhry123"}

	t.Run("Empty password.", func(t *testing.T) {
		usecase := New(repo)
		result := usecase.LoginAuthorized("fakhry@mail.id", "")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})

	t.Run("Empty email.", func(t *testing.T) {
		usecase := New(repo)
		result := usecase.LoginAuthorized("", "fakhry123")
		assert.Equal(t, "please input email and password", result)
		repo.AssertExpectations(t)
	})

	t.Run("Email not found", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything).Return(auth.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result := usecase.LoginAuthorized("ridho@mail.uk", "888ridho")
		assert.Equal(t, "email not found", result)
		repo.AssertExpectations(t)
	})

	t.Run("Wrong Password.", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything).Return(auth.Core{}, nil).Once()

		usecase := New(repo)
		result := usecase.LoginAuthorized("fakhry@mail.id", "fakhry123")
		assert.Equal(t, "wrong password", result)
		repo.AssertExpectations(t)
	})
}
