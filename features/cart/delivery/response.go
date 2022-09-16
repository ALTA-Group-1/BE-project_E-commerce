package delivery

import (
	"project/e-commerce/features/cart"
)

type CartResponse struct {
	ID       uint   `json:"id"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
	Images   string `json:"images,omitempty"`
	Price    int    `json:"price,omitempty"`
}

type TotalRes struct {
	TotalQuantity int
	Shipping      string
	TotalPrice    int
}

func fromCore(data cart.Core) CartResponse {
	return CartResponse{
		ID:       data.ID,
		Quantity: data.Quantity,
		Name:     data.ProductName,
		Images:   data.ProductImages,
		Price:    data.ProductPrice,
	}
}

func fromCoreList(data []cart.Core) ([]CartResponse, TotalRes) {

	var dataRes []CartResponse
	var totalRes TotalRes
	for i, v := range data {

		dataRes = append(dataRes, fromCore(v))
		totalRes.TotalQuantity += v.Quantity
		totalRes.Shipping = "Free"
		totalRes.TotalPrice += data[i].Quantity * data[i].ProductPrice

	}

	return dataRes, totalRes

}
