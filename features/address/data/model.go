package data

import "gorm.io/gorm"

type Core struct {
	TransactionID uint `gorm:"primary_key;ForeignKey:TransactionID"`
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type Transaction struct {
	gorm.Model
	Quantity    uint
	TotalPrice  uint
	OrderStatus string
	CartID      uint
}
