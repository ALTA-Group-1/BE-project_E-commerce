package data

import (
	"project/e-commerce/features/cart"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.DataInterface {
	return &cartData{
		db: db,
	}
}

func (repo *cartData) DeleteData(userID, productID int) (int, error) {
	var deleteData Cart
	tx := repo.db.Where("userID = ? AND productID = ?", userID, productID).Delete(&deleteData)
	return int(tx.RowsAffected), tx.Error
}
