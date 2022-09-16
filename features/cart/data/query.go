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

	var dbCek Cart
	repo.db.First(&dbCek, "product_id = ? AND user_id = ? ", data.ProductID, data.UserID)

	if dbCek.ID != 0 {
		return 2, nil
	}

	var dbCekProduct Product
	repo.db.First(&dbCekProduct, "id = ? AND user_id = ? ", data.ProductID, data.UserID)

	if dbCekProduct.ID > 0 {
		return 3, nil
	}

	dataModel := fromCore(data)

	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *cartData) DeleteData(userID, cartID int) (int, error) {

	txDel := repo.db.Unscoped().Where("user_id = ? AND id = ? ", userID, cartID).Delete(&Cart{})
	if txDel.Error != nil {
		return -1, txDel.Error
	}
	return int(txDel.RowsAffected), nil
}

func (repo *cartData) SelectByToken(token int) ([]cart.Core, error) {
	var dataCartCek []Results
	tx := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.name, products.images, products.price, carts.user_id, carts.product_id, carts.deleted_at").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&dataCartCek)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return toCoreList(dataCartCek), nil
}

func (repo *cartData) UpdateData(cartID, token int, update string) (int, error) {

	var qty int
	tx := repo.db.Model(&Cart{}).Select("quantity").Where("user_id = ? AND id = ? ", token, cartID).Scan(&qty)
	if tx.Error != nil {
		return -1, tx.Error
	}

	if update == "increment" {
		qty += 1
		txInc := repo.db.Model(&Cart{}).Where("user_id = ? AND id = ? ", token, cartID).Update("quantity", qty)
		if txInc.Error != nil {
			return -1, txInc.Error
		}

		return 1, nil

	} else {
		qty -= 1
		txDec := repo.db.Model(&Cart{}).Where("user_id = ? AND id = ? ", token, cartID).Update("quantity", qty)
		if txDec.Error != nil {
			return -1, txDec.Error
		}

		return 1, nil
	}

}
