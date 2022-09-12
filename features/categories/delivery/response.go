package delivery

import "project/e-commerce/features/categories"

type CategoriesResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func fromCore(data categories.Core) CategoriesResponse {
	return CategoriesResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}
