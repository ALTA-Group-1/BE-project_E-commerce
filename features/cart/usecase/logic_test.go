package usecase

import (
	"errors"
	"project/e-commerce/features/cart"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.CartData)
	dataInput := cart.Core{ProductID: 1}

	t.Run("Empty data.", func(t *testing.T) {
		dataInput := cart.Core{ProductID: 0}

		usecase := New(repo)
		result, err := usecase.PostData(dataInput)
		assert.EqualError(t, err, "data tidak boleh kosong")
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Insert data.", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.PostData(dataInput)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Success Insert data.", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		result, err := usecase.PostData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := new(mocks.CartData)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteCart(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteCart(1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetByToken(t *testing.T) {
	repo := new(mocks.CartData)

	dataInput := []cart.Core{{ID: 1, ProductName: "Sneakers", ProductImages: "https://images.jpg", ProductPrice: 200000, Quantity: 1, UserID: 1}}

	t.Run("Success Get data by Token.", func(t *testing.T) {
		repo.On("SelectByToken", mock.Anything).Return(dataInput, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetByToken(1)
		assert.NoError(t, err)
		assert.Equal(t, dataInput, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data by Token.", func(t *testing.T) {
		repo.On("SelectByToken", mock.Anything).Return([]cart.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetByToken(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestPutData(t *testing.T) {
	repo := new(mocks.CartData)

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutData(1, 1, "increment")
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutData(1, 1, "decrement")
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}
