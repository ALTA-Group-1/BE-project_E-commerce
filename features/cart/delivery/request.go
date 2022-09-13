package delivery

import (
	"project/e-commerce/features/cart"
)

type CartRequest struct {
	Quantity int `json:"quantity" form:"quantity"`
}

func toCore(data CartRequest) cart.Core {
	return cart.Core{
		Quantity: data.Quantity,
	}
}
