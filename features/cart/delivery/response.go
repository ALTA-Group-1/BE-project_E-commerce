package delivery

import (
	"project/e-commerce/features/cart"
)

type CartResponse struct {
	ID        uint   `json:"id"`
	Quantity  int    `json:"quantity"`
	Name      string `json:"name"`
	Images    string `json:"images,omitempty"`
	Price     int    `json:"price,omitempty"`
	UserID    int
	ProductID int `json:"product_id,omitempty"`
}

func fromCore(data cart.Core) CartResponse {
	return CartResponse{
		ID:        data.ID,
		Quantity:  data.Quantity,
		Name:      data.ProductName,
		Images:    data.ProductImages,
		Price:     data.ProductPrice,
		UserID:    data.UserID,
		ProductID: int(data.ProductID),
	}
}

func fromCoreList(data []cart.Core) []CartResponse {

	var dataRes []CartResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes

}
