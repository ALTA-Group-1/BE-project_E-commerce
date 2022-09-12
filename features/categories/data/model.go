package data

import (
	"project/e-commerce/features/categories"
	productModel "project/e-commerce/features/product/data"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name    string
	Product productModel.Product
}

func fromCore(dataCore categories.Core) Categories {
	dataModel := Categories{
		Name: dataCore.Name,
	}
	return dataModel
}
