package delivery

import (
	"project/e-commerce/features/product"
)

type ProductResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Images string `json:"images,omitempty"`
	Price  int    `json:"price,omitempty"`
	Stock  int    `json:"stock,omitempty"`
	Desc   string `json:"desc,omitempty"`
}

type ProductResponseList struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Images string `json:"images,omitempty"`
	Price  int    `json:"price,omitempty"`
}

func fromCore(data product.Core) ProductResponse {
	return ProductResponse{
		ID:     data.ID,
		Name:   data.Name,
		Images: data.Images,
		Price:  data.Price,
		Stock:  data.Stock,
		Desc:   data.Desc,
	}
}

func fromCoreListMyProduct(data []product.Core) []ProductResponse {

	var dataRes []ProductResponse
	for _, v := range data {
		dataRes = append(dataRes, fromCore(v))
	}

	return dataRes

}

func fromCoreList(data []product.Core) []ProductResponseList {

	var dataRes []ProductResponseList
	for _, v := range data {
		dataRes = append(dataRes, ProductResponseList{
			ID:     v.ID,
			Images: v.Images,
			Name:   v.Name,
			Price:  v.Price,
		})
	}

	return dataRes

}
