package delivery

import (
	"project/e-commerce/features/product"
)

type ProductDelivery struct {
	productUsecase product.UsecaseInterface
}

// func New(e *echo.Echo, usecase user.UsecaseInterface) {
// 	handler := &ProductDelivery{
// 		userUsecase: usecase,
// 	}

// }
