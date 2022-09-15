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
	Desc         string
	UserID       uint
	CategoriesID uint
	CartID       []Cart `gorm:"foreignKey:ProductID"`
}

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	UserID    uint
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
		Name:         dataCore.Name,
		Images:       dataCore.Images,
		Price:        dataCore.Price,
		Desc:         dataCore.Desc,
		UserID:       uint(dataCore.UserID),
		CategoriesID: uint(dataCore.CategoriesID),
	}
	return productModel
}

func (data *Product) toCore() product.Core {
	return product.Core{
		ID:           data.ID,
		Name:         data.Name,
		Images:       data.Images,
		Price:        data.Price,
		Desc:         data.Desc,
		UserID:       int(data.UserID),
		CategoriesID: int(data.CategoriesID),
	}
}

func toCoreList(data []Product) []product.Core {
	var dataCore []product.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
