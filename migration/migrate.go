package migration

import (
	productModel "project/e-commerce/features/product/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&productModel.Product{})

}
