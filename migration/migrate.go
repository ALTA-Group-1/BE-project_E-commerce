package migration

import (
	cartModel "project/e-commerce/features/cart/data"
	categoryModel "project/e-commerce/features/categories/data"
	productModel "project/e-commerce/features/product/data"
	transactionModel "project/e-commerce/features/transaction/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&productModel.Product{})
	db.AutoMigrate(&categoryModel.Categories{})
	db.AutoMigrate(&cartModel.Cart{})
	db.AutoMigrate(&transactionModel.Transaction{})
}
