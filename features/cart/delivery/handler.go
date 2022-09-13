package delivery

import (
	"project/e-commerce/features/cart"
)

type CartDelivery struct {
	cartUsecase cart.UsecaseInterface
}

// func New(e *echo.Echo, usecase cart.UsecaseInterface) {
// 	handler := &CartDelivery{
// 		cartUsecase: usecase,
// 	}

// }
