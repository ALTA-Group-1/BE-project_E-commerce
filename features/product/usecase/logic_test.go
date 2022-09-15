package usecase

import (
	"errors"
	"project/e-commerce/features/product"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.ProductData)
	dataInput := product.Core{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Desc: "Size L, Original verified, Made in US", CategoriesID: 3}

	t.Run("Success Insert data.", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		result, err := usecase.PostData(dataInput)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
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

	t.Run("Empty data.", func(t *testing.T) {
		dataInput := product.Core{Name: "", Images: "", Price: 0, Desc: "", CategoriesID: 0}

		usecase := New(repo)
		result, err := usecase.PostData(dataInput)
		assert.EqualError(t, err, "data tidak boleh kosong")
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetAllProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	dataProduct := []product.Core{{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Desc: "Size L, Original verified, Made in US", CategoriesID: 3}}

	t.Run("Success Get all data.", func(t *testing.T) {
		repo.On("SelectAllProduct", mock.Anything, mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetAllProduct(1, "Style")
		assert.NoError(t, err)
		assert.Equal(t, dataProduct, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get all data.", func(t *testing.T) {
		repo.On("SelectAllProduct", mock.Anything, mock.Anything).Return([]product.Core{}, errors.New("error")).Once()

		usecase := New(repo)
		result, err := usecase.GetAllProduct(1, "Style")
		assert.Error(t, err)
		assert.Equal(t, []product.Core([]product.Core(nil)), result)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ProductData)
	dataProduct := product.Core{ID: 1, Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Desc: "Size L, Original verified, Made in US", CategoriesID: 3}

	t.Run("Success Get data by Id.", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, dataProduct, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get data by Id.", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetById(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}

func TestPutData(t *testing.T) {
	repo := new(mocks.ProductData)
	newData := product.Core{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Desc: "Size L, Original verified, Made in US", CategoriesID: 3}

	t.Run("Success Update data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.PutData(1, newData)
		assert.NoError(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Update data", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.PutData(1, newData)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.ProductData)

	t.Run("Success Delete data.", func(t *testing.T) {
		repo.On("DeleteByToken", mock.Anything, mock.Anything).Return(1, nil).Once()

		usecase := New(repo)

		result, err := usecase.DeleteData(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, 1, result)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete data", func(t *testing.T) {
		repo.On("DeleteByToken", mock.Anything, mock.Anything).Return(-1, errors.New("Error")).Once()

		usecase := New(repo)

		result, err := usecase.DeleteData(1, 1)
		assert.Error(t, err)
		assert.Equal(t, -1, result)
		repo.AssertExpectations(t)
	})
}

func TestGetMyProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	dataProduct := []product.Core{{Name: "Converse Allstar", Images: "https://cf.shopee.co.id/file/6614946421f206cff75a2f79310a2e35", Price: 1000000, Desc: "Size L, Original verified, Made in US", UserID: 1, CategoriesID: 3}}

	t.Run("Success Get my product.", func(t *testing.T) {
		repo.On("SelectMyProduct", mock.Anything).Return(dataProduct, nil).Once()

		usecase := New(repo)
		result, err := usecase.GetMyProduct(1)
		assert.NoError(t, err)
		assert.Equal(t, result[0].UserID, dataProduct[0].UserID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get my product.", func(t *testing.T) {
		repo.On("SelectMyProduct", mock.Anything).Return([]product.Core{}, errors.New("Error")).Once()

		usecase := New(repo)
		result, err := usecase.GetMyProduct(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, result)
		repo.AssertExpectations(t)
	})
}
