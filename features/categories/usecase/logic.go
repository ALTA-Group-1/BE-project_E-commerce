package usecase

import (
	"project/e-commerce/features/categories"
)

type categoriesUsecase struct {
	categoriesData categories.DataInterface
}

func New(data categories.DataInterface) categories.UsecaseInterface {
	return &categoriesUsecase{
		categoriesData: data,
	}
}
