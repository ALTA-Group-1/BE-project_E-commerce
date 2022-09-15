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
	dataInput := auth.Core{ID: 1, Email: "fakhry@mail.id", Password: "$2a$10$3qSIi7BiTknraN3A9tRX/eoI4N9yuln/oWI8Ft9KcrZNF3ec6jIHK"}

	t.Run("success password.", func(t *testing.T) {
		repo.On("LoginUser", mock.Anything, mock.Anything).Return(dataInput, nil).Once()

		usecase := New(repo)
		result := usecase.LoginAuthorized("fakhry@mail.id", "12345")
		assert.NotEqual(t, "please input email and password", result)
		assert.NotEqual(t, "email not found", result)
		assert.NotEqual(t, "wrong password", result)
		assert.NotEqual(t, "error to created token", result)
		repo.AssertExpectations(t)
	})

	t.Run("success login", func(t *testing.T) {
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
