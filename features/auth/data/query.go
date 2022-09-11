package data

import (
	"project/e-commerce/features/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.DataInterface {
	return &authData{
		db: db,
	}
}
