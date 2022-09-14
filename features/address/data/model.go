package data

import (
	"project/e-commerce/features/address"

	"gorm.io/gorm"
)

type Address struct {
	TransactionID uint `gorm:"primary_key;ForeignKey:TransactionID"`
	Street        string
	City          string
	Province      string
	PostCode      uint
}

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

type DBTransaction struct {
	Cart_id  uint
	Quantity int
	Stock    int
	ID       uint
}

func toDb(data address.Core) Address {
	return Address{
		TransactionID: data.TransactionID,
		Street:        data.Street,
		City:          data.City,
		Province:      data.Province,
		PostCode:      data.PostCode,
	}

}
