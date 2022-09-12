package delivery

import (
	"project/e-commerce/features/categories"
)

type CategoriesRequest struct {
	Name string `json:"name" form:"name"`
}

func toCore(data CategoriesRequest) categories.Core {
	return categories.Core{
		Name: data.Name,
	}
}
