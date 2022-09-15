package usecase

import (
	"errors"
	"project/e-commerce/features/transaction"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.TransactionData)
	// mockOrder := transaction.Core{Quantity: 10, TotalPrice: 100000, OrderStatus: "waiting", CartID: 12}
	address := transaction.AddressCore{Street: "jl. kusuma", City: "bandung", Province: "jawa barat", PostCode: 17654}
	payment := transaction.PaymentCore{Visa: "mega bank", Name: "dian", Number: 836502648396, Cvv2: 9248, Month: 3, Year: 2025}

	t.Run("order success", func(t *testing.T) {

		repo.On("InsertData", mock.Anything, mock.Anything, mock.Anything).Return(1, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.PostData(1, address, payment)
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("empty data", func(t *testing.T) {

		address.City = ""
		address.PostCode = 0
		address.Province = ""
		address.Street = ""
		address.TransactionID = 0

		payment.TransactionID = 0
		payment.Visa = ""
		payment.Name = ""
		payment.Number = 0
		payment.Cvv2 = 0
		payment.Month = 0
		payment.Year = 0

		useCase := New(repo)
		res, err := useCase.PostData(1, address, payment)
		assert.Error(t, err, "data tidak boleh kosong")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteOrder(t *testing.T) {
	repo := new(mocks.TransactionData)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("CancelOrder", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteOrder(1, "waiting")
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("CancelOrder", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteOrder(1, "waiting")
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetOrder(t *testing.T) {
	repo := new(mocks.TransactionData)
	dataOrder := []transaction.HistoryOrder{{Images: "https://images.jpg", Name: "Sneakers", Price: 200000, Quantity: 2}}

	t.Run("Success Get order.", func(t *testing.T) {
		repo.On("SelectOrder", mock.Anything).Return(dataOrder, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetOrder(1)
		assert.NoError(t, err)
		assert.Equal(t, dataOrder, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get order.", func(t *testing.T) {
		repo.On("SelectOrder", mock.Anything).Return([]transaction.HistoryOrder{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetOrder(1)
		assert.Error(t, err)
		assert.Equal(t, []transaction.HistoryOrder([]transaction.HistoryOrder(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestPutStatus(t *testing.T) {
	repo := new(mocks.TransactionData)

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateStatus", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutStatus(1, "waiting")
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateStatus", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutStatus(1, "waiting")
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}
