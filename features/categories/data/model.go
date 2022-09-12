package data

import (
	productModel "project/e-commerce/features/product/data"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name    string
	Product productModel.Product
}
