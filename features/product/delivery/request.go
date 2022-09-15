package delivery

import (
	"project/e-commerce/features/product"
)

type ProductRequest struct {
	Name       string `json:"name" form:"name"`
	Images     string `json:"images" form:"images"`
	Price      int    `json:"price" form:"price"`
	Desc       string `json:"desc" form:"desc"`
	Categories int    `json:"categories,omitempty"`
	UserID     int
}

func toCore(data ProductRequest) product.Core {
	return product.Core{
		Name:         data.Name,
		Images:       data.Images,
		Price:        data.Price,
		Desc:         data.Desc,
		CategoriesID: data.Categories,
		UserID:       data.UserID,
	}
}
