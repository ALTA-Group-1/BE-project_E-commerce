package data

import (
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
