package delivery

import (
	"project/e-commerce/features/categories"
)

type CategoriesResponse struct {
	ID        uint `json:"id"`
	ProductID uint
	Images    string
	Name      string
	Price     int
}

func fromCore(data categories.Core) CategoriesResponse {
	return CategoriesResponse{
		ID:        data.ID,
		ProductID: data.ProductID,
		Images:    data.ProductImages,
		Name:      data.ProductName,
		Price:     data.ProductPrice,
	}
}

func fromCoreList(data []categories.Core) []CategoriesResponse {
	var dataResponse []CategoriesResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
