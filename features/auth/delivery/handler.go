package delivery

import (
	"project/e-commerce/features/auth"
)

type AuthDelivery struct {
	authUsecase auth.UsecaseInterface
}

// func New(e *echo.Echo, usecase auth.UsecaseInterface) {
// 	handler := &AuthDelivery{
// 		authUsecase: usecase,
// 	}

// }
