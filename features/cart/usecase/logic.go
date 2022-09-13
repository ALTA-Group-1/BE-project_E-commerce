package usecase

import (
	"project/e-commerce/features/cart"
)

type cartUsecase struct {
	cartData cart.DataInterface
}

func New(data cart.DataInterface) cart.UsecaseInterface {
	return &cartUsecase{
		cartData: data,
	}
}
