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
