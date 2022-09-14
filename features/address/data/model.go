package data

import "gorm.io/gorm"

type Core struct {
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
