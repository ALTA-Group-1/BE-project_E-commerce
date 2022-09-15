package usecase

import (
	"errors"
	"project/e-commerce/features/address"
)

type addressUsecase struct {
	addressData address.DataInterface
}

func New(data address.DataInterface) address.UsecaseInterface {
	return &addressUsecase{
		addressData: data,
	}
}

func (usecase *addressUsecase) PostData(token int, data address.Core) (int, error) {
	if data.City == "" || data.Province == "" || data.Street == "" || data.PostCode == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.addressData.InsertData(token, data)
	if err != nil {
		return -1, err
	}

	return row, nil

}
