package data

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string
	Images       string
	Price        int
	Stock        int
	Desc         string
	UserID       uint
	CategoriesID uint
}

type Categories struct {
	gorm.Model
	Name    string
	Product []Product
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Product  []Product
}

func fromCore(dataCore product.Core) Product {
	productModel := Product{
		Name:   dataCore.Name,
		Images: dataCore.Images,
		Price:  dataCore.Price,
		Stock:  dataCore.Stock,
		Desc:   dataCore.Desc,
	}
	return productModel
}

func (data *Product) toCore() product.Core {
	return product.Core{
		ID:     data.ID,
		Name:   data.Name,
		Images: data.Images,
		Price:  data.Price,
		Stock:  data.Stock,
		Desc:   data.Desc,
	}
}

func toCoreList(data []Product) []product.Core {
	var dataCore []product.Core
	for key, _ := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
