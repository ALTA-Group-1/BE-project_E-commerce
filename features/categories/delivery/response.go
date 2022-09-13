package delivery

import (
	"project/e-commerce/features/categories"
)

type CategoriesResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	ProductID     uint   `json:"productid,omitempty"`
	ProductName   string `json:"productname,omitempty"`
	ProductImages string `json:"productimages,omitempty"`
	ProductPrice  int    `json:"productprice,omitempty"`
	ProductStock  int    `json:"productstock,omitempty"`
	ProductDesc   string `json:"productdesc,omitempty"`
}

func fromCore(data categories.Core) CategoriesResponse {
	return CategoriesResponse{
		ID:            data.ID,
		Name:          data.Name,
		ProductID:     data.ProductID,
		ProductName:   data.ProductName,
		ProductImages: data.ProductImages,
		ProductPrice:  data.ProductPrice,
		ProductStock:  data.ProductStock,
		ProductDesc:   data.ProductDesc,
	}
}

func fromCoreList(data []categories.Core) []CategoriesResponse {
	var dataResponse []CategoriesResponse
	for _, v := range data {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
