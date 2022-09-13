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
		Quantity:  dataCore.Quantity,
		ProductID: dataCore.ProductID,
		UserID:    uint(dataCore.UserID),
	}

	return productModel

}

func toCore(dataCart Cart, dataProduct Product) cart.Core {

	return cart.Core{
		ID:            dataCart.ID,
		ProductID:     dataCart.ProductID,
		ProductImages: dataProduct.Images,
		ProductName:   dataProduct.Name,
		ProductPrice:  dataProduct.Price,
	}

}

func toCoreList(dataCart []Cart, dataProduct []Product) []cart.Core {
	var dataCore []cart.Core
	for _, v := range dataCart {
		for _, v1 := range dataProduct {
			if v.ProductID == v1.ID {
				dataCore = append(dataCore, toCore(v, v1))
			}
		}
	}

	return dataCore

}
