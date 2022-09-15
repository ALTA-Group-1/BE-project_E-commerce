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

type Results struct {
	ID         uint
	Quantity   int
	Name       string
	Images     string
	Price      int
	UserID     uint
	ProductID  uint
	Deleted_at string
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

func (res *Results) toCore() cart.Core {

	return cart.Core{
		ID:            res.ID,
		Quantity:      res.Quantity,
		ProductID:     uint(res.ProductID),
		ProductImages: res.Images,
		ProductName:   res.Name,
		ProductPrice:  res.Price,
		UserID:        int(res.UserID),
	}

}

func toCoreList(dataCart []Results) []cart.Core {
	var dataCore []cart.Core
	for key := range dataCart {
		if dataCart[key].Deleted_at == "" {
			dataCore = append(dataCore, dataCart[key].toCore())
		}
	}

	return dataCore

}
