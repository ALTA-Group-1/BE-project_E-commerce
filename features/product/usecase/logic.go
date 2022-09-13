package usecase

import (
	"errors"
	"project/e-commerce/features/product"
)

type productUsecase struct {
	productData product.DataInterface
}

func New(data product.DataInterface) product.UsecaseInterface {
	return &productUsecase{
		productData: data,
	}
}

func (usecase *productUsecase) PostData(data product.Core) (int, error) {
	if data.Name == "" || data.Images == "" || data.Price == 0 || data.Stock == 0 || data.Desc == "" || data.CategoriesID == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.productData.InsertData(data)
	return row, err
}

func (usecase *productUsecase) DeleteData(token int) (int, error) {
	row, err := usecase.productData.DeleteByToken(token)

	return row, err
}
