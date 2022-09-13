package usecase

import (
	"errors"
	"project/e-commerce/features/user"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userData user.DataInterface
}

func New(data user.DataInterface) user.UsecaseInterface {
	return &userUsecase{
		userData: data,
	}
}

func (usecase *userUsecase) PostData(data user.Core) (int, error) {
	if data.Name == "" || data.Email == "" || data.Password == "" || data.Phone == "" || data.Address == "" {
		return -1, errors.New("data tidak boleh kosong")
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	data.Password = string(hashPass)
	row, err := usecase.userData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (service *userUsecase) GetByToken(token int) (user.Core, error) {
	dataId, err := service.userData.SelectByToken(token)
	if err != nil {
		return user.Core{}, err
	}
	return dataId, nil
}

func (usecase *userUsecase) PutData(newData user.Core) (int, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(newData.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}
	newData.Password = string(hashPass)
	row, err := usecase.userData.UpdateData(newData)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *userUsecase) DeleteData(token int) (int, error) {
	row, err := usecase.userData.DeleteByToken(token)
	if err != nil {
		return -1, err
	}
	return row, nil
}
