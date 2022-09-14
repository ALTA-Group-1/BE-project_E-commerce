package delivery

import (
	"project/e-commerce/features/cart"
)

type CartRequest struct {
	Quantity  int `json:"quantity" form:"quantity"`
	ProductID int `json:"productID" form:"productID"`
	UserID    int
}

func toCore(data CartRequest) cart.Core {
	return cart.Core{
		Quantity:  data.Quantity,
		ProductID: uint(data.ProductID),
		UserID:    data.UserID,
	}
}
