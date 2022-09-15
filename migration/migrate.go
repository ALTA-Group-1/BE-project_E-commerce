package migration

import (
	cartModel "project/e-commerce/features/cart/data"
	productModel "project/e-commerce/features/product/data"
	transactionModel "project/e-commerce/features/transaction/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&productModel.Product{})
	db.AutoMigrate(&productModel.Categories{})
	db.AutoMigrate(&cartModel.Cart{})
	db.AutoMigrate(&transactionModel.Transaction{})
	db.AutoMigrate(&transactionModel.Address{})
	db.AutoMigrate(&transactionModel.Payment{})
}
