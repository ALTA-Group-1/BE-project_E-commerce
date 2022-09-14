package data

import (
	"errors"
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
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.images, products.name, products.price, users.id, products.id").Joins("left join carts on carts.products_id = products.id").Where("products.user_id = ?", token).Scan(&dataCart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(dataCart), nil
}

func (repo *cartData) UpdatePlusData(cartID int, increment string) (int, error) {
	var dataProduct cart.Core
	dataModel := fromCore(dataProduct)

	tx := repo.db.Raw("UPDATE carts SET quantity = (? + 1) WHERE carts_id = ? AND products_id = ?", dataModel.Quantity, cartID, dataModel.ProductID).Scan(&cartID)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil
}

func (repo *cartData) UpdateMinusData(cartID int, decrement string) (int, error) {
	var dataProduct cart.Core
	dataModel := fromCore(dataProduct)

	tx := repo.db.Raw("UPDATE carts SET quantity = (? - 1) WHERE carts_id = ? AND products_id = ?", dataModel.Quantity, cartID, dataModel.ProductID).Scan(&cartID)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil
}
