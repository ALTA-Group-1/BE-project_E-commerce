package migration

import (
	categoriesModel "project/e-commerce/features/categories/data"
	userModel "project/e-commerce/features/user/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&categoriesModel.Categories{})

}
