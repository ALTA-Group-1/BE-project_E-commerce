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

func (data *Categories) toCore() categories.Core {
	return categories.Core{
		ID:   data.ID,
		Name: data.Name,
	}
}

func toCoreList(data []Categories) []categories.Core {
	var dataCore []categories.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
