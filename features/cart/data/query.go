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

func (repo *cartData) InsertData(data cart.Core) (int, error) {
	dataModel := fromCore(data)

	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *cartData) DeleteData(userID, cartID int) (int, error) {
	var deleteData Cart
	tx := repo.db.Where("user_id = ? AND id = ?", userID, cartID).Delete(&deleteData)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *cartData) SelectByToken(token int) ([]cart.Core, error) {
	var dataCart []Results
	// tx := repo.db.Model(&Cart{}).Where("user_id = ?", token).Find(&dataCart)
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.name, products.images, products.price, carts.user_id, carts.product_id").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&dataCart)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(dataCart), nil
}
