package usecase

import (
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

func (usecase *productUsecase) DeleteData(token int) (int, error) {
	row, err := usecase.productData.DeleteByToken(token)
	return row, err
}
