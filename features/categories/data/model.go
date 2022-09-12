package data

import (
	"project/e-commerce/features/categories"
	productModel "project/e-commerce/features/product/data"
	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name    string
	Product []Product `gorm:"foreignKey:CategoriesID"`
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
