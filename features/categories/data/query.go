package data

import (
	"project/e-commerce/features/categories"

	"gorm.io/gorm"
)

type categoriesData struct {
	db *gorm.DB
}

func New(db *gorm.DB) categories.DataInterface {
	return &categoriesData{
		db: db,
	}
}
