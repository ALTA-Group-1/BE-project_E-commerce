package data

import (
	"project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	UserID    uint
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
	CartID       []Cart
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Product  []Product
	Cart     []Cart
}

func fromCore(dataCore cart.Core) Cart {
	productModel := Cart{
		Quantity: dataCore.Quantity,
	}
	return productModel
}

func (data *Cart) toCore() cart.Core {
	return cart.Core{
		ID:       data.ID,
		Quantity: data.Quantity,
	}
}

func toCoreList(data []Cart) []cart.Core {
	var dataCore []cart.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
