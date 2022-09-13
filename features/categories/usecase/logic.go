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

func (usecase *categoriesUsecase) GetAll(id int) ([]categories.Core, error) {
	results, err := usecase.categoriesData.GetAllData(id)
	if err != nil {
		return nil, err
	}
	return results, nil
}
