package usecase

import (
	"errors"
	"project/e-commerce/features/user"
	"project/e-commerce/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPostData(t *testing.T) {
	repo := new(mocks.DataInterface)
	input := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta"}

	t.Run("create success", func(t *testing.T) {

		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("error create data", func(t *testing.T) {

		repo.On("InsertData", mock.Anything).Return(-1, errors.New("there is some error")).Once()

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.EqualError(t, err, "there is some error")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	t.Run("empty data", func(t *testing.T) {

		input.Name = ""
		input.Password = ""
		input.Address = ""
		input.Email = ""
		input.Phone = ""

		useCase := New(repo)
		res, err := useCase.PostData(input)
		assert.EqualError(t, err, "data tidak boleh kosong")
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

}

func TestGetByToken(t *testing.T) {
	repo := new(mocks.DataInterface)
	data := user.Core{Name: "coco", Email: "coco@gmail.com", Phone: "0812345", Address: "jakarta"}

	t.Run("get data success", func(t *testing.T) {

		repo.On("SelectByToken", mock.Anything).Return(data, nil).Once()

		useCase := New(repo)

		res, err := useCase.GetByToken(1)
		assert.NoError(t, err)
		assert.Equal(t, data, res)
		repo.AssertExpectations(t)
	})

	t.Run("error get data", func(t *testing.T) {

		repo.On("SelectByToken", mock.Anything).Return(user.Core{}, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.GetByToken(1)
		assert.Error(t, err)
		assert.NotEqual(t, 1, res)
		repo.AssertExpectations(t)
	})
}

func TestPutData(t *testing.T) {
	repo := new(mocks.DataInterface)
	newData := user.Core{Name: "coco", Email: "coco@gmail.com", Password: "123", Phone: "0812345", Address: "jakarta"}

	t.Run("success update data", func(t *testing.T) {

		repo.On("UpdateData", mock.Anything).Return(-1, nil).Once()

		useCase := New(repo)

		res, err := useCase.PutData(newData)
		assert.NoError(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})

	// t.Run("Generate Hash Error", func(t *testing.T) {

	// 	newData.Password = "coco123"

	// 	repo.On("UpdateData", mock.Anything).Return(-1, errors.New("there is some error")).Once()

	// 	useCase := New(repo)

	// 	res, _ := useCase.PutData(newData)
	// 	assert.Equal(t, -1, res)
	// 	repo.AssertExpectations(t)
	// })
}

func TestDeleteData(t *testing.T) {
	repo := new(mocks.DataInterface)

	t.Run("success delete data", func(t *testing.T) {

		repo.On("DeleteByToken", mock.Anything).Return(-1, nil).Once()

		useCase := New(repo)

		delete, err := useCase.DeleteData(1)
		assert.Nil(t, err)
		assert.Equal(t, -1, delete)
		repo.AssertExpectations(t)
	})

	t.Run("error delete data", func(t *testing.T) {

		repo.On("DeleteByToken", mock.Anything).Return(-1, errors.New("error")).Once()

		useCase := New(repo)

		res, err := useCase.DeleteData(1)
		assert.Error(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}
