package delivery

import (
	"project/e-commerce/features/categories"
)

type CategoriesDelivery struct {
	categoriesUsecase categories.UsecaseInterface
}

// func New(e *echo.Echo, usecase categories.UsecaseInterface) {
// 	handler := &CategoriesDelivery{
// 		categoriesUsecase: usecase,
// 	}

// }
