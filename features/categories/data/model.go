package data

import (
	"project/e-commerce/features/categories"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name    string
	Product []Product `gorm:"foreignKey:CategoriesID"`
}

type Results struct {
	ProductID uint
	Images    string
	Name      string
	Price     int
}

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

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Product  []Product
}

func fromCore(dataCore categories.Core) Categories {
	dataModel := Categories{
		Name: dataCore.Name,
	}
	return dataModel
}

func (res *Results) toCore() categories.Core {
	return categories.Core{
		ProductID:     res.ProductID,
		ProductImages: res.Images,
		ProductName:   res.Name,
		ProductPrice:  res.Price,
	}
}

func toCoreList(data []Results) []categories.Core {
	var dataCore []categories.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
