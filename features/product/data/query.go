package data

import (
	"project/e-commerce/features/product"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.DataInterface {
	return &productData{
		db: db,
	}
}

func (repo *productData) DeleteByToken(token int) (int, error) {
	var deleteData User
	tx := repo.db.Delete(&deleteData, token)
	return int(tx.RowsAffected), tx.Error
}
