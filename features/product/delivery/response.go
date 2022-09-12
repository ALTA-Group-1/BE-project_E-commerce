package delivery

import (
	"project/e-commerce/features/product"
)

type ProductResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Images     string `json:"images,omitempty"`
	Price      int    `json:"price,omitempty"`
	Stock      int    `json:"stock,omitempty"`
	Desc       string `json:"desc,omitempty"`
	User       string `json:"user,omitempty"`
	Categories string `json:"categories,omitempty"`
}

func fromCore(data product.Core) ProductResponse {
	return ProductResponse{
		ID:     data.ID,
		Name:   data.Name,
		Images: data.Images,
		Price:  data.Price,
		Stock:  data.Stock,
		Desc:   data.Desc,
		User: data.User{
			Name: Core.Name,
		},
		Categories: data.Categories{
			Categories: Core.Name,
		},
	}
}
