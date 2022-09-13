package delivery

import (
	"project/e-commerce/features/cart"
)

type CartRequest struct {
	Quantity  int `json:"quantity" form:"quantity"`
	ProductID int `json:"productID" form:"productID"`
	UserID    int `json:"userID" form:"userID"`
}

func toCore(data CartRequest) cart.Core {
	return cart.Core{
		Quantity:  data.Quantity,
		ProductID: uint(data.ProductID),
		UserID:    data.UserID,
	}
}
