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

func (repo *cartData) DeleteData(userID, cartID int) (int, error) {
	var deleteData Cart
	tx := repo.db.Where("user_id = ? AND id = ?", userID, cartID).Delete(&deleteData)
	return int(tx.RowsAffected), tx.Error
}

func (repo *cartData) SelectByToken(token int) ([]cart.Core, error) {

	var dataCart []Results
	// tx := repo.db.Model(&Cart{}).Where("user_id = ?", token).Find(&dataCart)
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.images, products.name, products.price, users.id, products.id").Joins("left join carts on carts.products_id = products.id").Scan(&dataCart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(dataCart), nil
}
