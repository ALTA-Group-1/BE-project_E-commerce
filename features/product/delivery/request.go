package delivery

import (
	"project/e-commerce/features/product"
)

type ProductRequest struct {
	Name       string `json:"name" form:"name"`
	Images     string `json:"images" form:"images"`
	Price      int    `json:"price" form:"price"`
	Stock      int    `json:"stock" form:"stock"`
	Desc       string `json:"desc" form:"desc"`
	Categories int    `json:"categories,omitempty"`
}

func toCore(data ProductRequest) product.Core {
	return product.Core{
		Name:         data.Name,
		Images:       data.Images,
		Price:        data.Price,
		Stock:        data.Stock,
		Desc:         data.Desc,
		CategoriesID: data.Categories,
	}
}
