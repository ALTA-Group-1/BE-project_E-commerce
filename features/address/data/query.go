package data

import (
	"project/e-commerce/features/address"

	"gorm.io/gorm"
)

type addressData struct {
	db *gorm.DB
}

func New(db *gorm.DB) address.DataInterface {
	return &addressData{
		db: db,
	}
}

func (repo *addressData) InsertData(token int, data address.Core) (int, error) {

	var id []int
	str := "waiting"
	tx := repo.db.Model(&Cart{}).Select("transactions.id").Joins("inner join transactions on transactions.cart_id = carts.id").Where("carts.user_id = ? AND transactions.order_status = ?", token, str).Scan(&id)
	if tx.Error != nil {
		return -1, tx.Error
	}

	for _, v := range id {
		data.TransactionID = uint(v)
		dataCreate := toDb(data)
		txCreate := repo.db.Create(&dataCreate)
		if txCreate.Error != nil {
			return -1, txCreate.Error
		}
	}

	var dbTransaction []DBTransaction
	txTransactions := repo.db.Model(&Product{}).Select("carts.id, carts.quantity, products.stock, products.id").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&dbTransaction)
	if txTransactions.Error != nil {
		return -1, txTransactions.Error
	}

	for _, value := range dbTransaction {
		txStock := repo.db.Model(&Product{}).Where("id = ?", value.ID).Update("stock", value.Stock-value.Quantity)
		if txStock.Error != nil {
			return -1, txStock.Error
		}
		txDel := repo.db.Where("id = ?", value.Cart_id).Delete(&Cart{})
		if txDel.Error != nil {
			return -1, txDel.Error
		}
	}

	for _, valueId := range id {
		txUpdate := repo.db.Model(&Product{}).Where("id = ?", valueId).Update("order_status", "confirm")
		if txUpdate.Error != nil {
			return -1, txUpdate.Error
		}
	}

	return -1, nil

}
