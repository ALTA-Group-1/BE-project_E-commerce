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

func (repo *cartData) InsertData(data cart.Core) (int, error) {

	var cek int
	txUpd := repo.db.Model(&Cart{}).Select("carts.quantity").Where("product_id = ? AND user_id = ? AND deleted_at = NULL", data.ProductID, data.UserID).Scan(&cek)
	if txUpd.Error != nil {
		return -1, txUpd.Error
	}

	if cek > 0 {
		return 2, nil
	}

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
	var dataCartCek []Results
	// tx := repo.db.Model(&Cart{}).Where("user_id = ?", token).Find(&dataCart)
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.name, products.images, products.price, carts.user_id, carts.product_id, carts.deleted_at").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&dataCartCek)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(dataCartCek), nil
}

func (repo *cartData) UpdatePlusData(cartID int, increment string) (int, error) {
	var dataProduct cart.Core
	dataModel := fromCore(dataProduct)

	if cartID != 0 && increment == "increment" {
		tx := repo.db.Raw("UPDATE carts SET quantity = (? + 1) WHERE carts_id = ?", dataModel.Quantity, cartID).Scan(&dataModel)
		if tx.Error != nil {
			return -1, tx.Error
		}
		if tx.RowsAffected == 0 {
			return -1, errors.New("failed to update quantity")
		}
	}

	return 1, nil
}

func (repo *cartData) UpdateMinusData(cartID int, decrement string) (int, error) {
	var dataProduct cart.Core
	dataModel := fromCore(dataProduct)

	if cartID != 0 && decrement == decrement {
		tx := repo.db.Raw("UPDATE carts SET quantity = (? + 1) WHERE carts_id = ?", dataModel.Quantity, cartID).Scan(&dataModel)
		if tx.Error != nil {
			return -1, tx.Error
		}
		if tx.RowsAffected == 0 {
			return -1, errors.New("failed to update quantity")
		}
	}

	return 1, nil
}
