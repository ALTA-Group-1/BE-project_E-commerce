package data

import (
	"log"
	"project/e-commerce/features/transaction"

	"gorm.io/gorm"
)

type transactionData struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.DataInterface {
	return &transactionData{
		db: db,
	}
}

func (repo *transactionData) InsertData(token int) (int, error) {

	var inputTransaction []result
	tx := repo.db.Model(&Product{}).Select("carts.quantity, products.price, carts.id").Joins("inner join carts on carts.product_id = products.id").Where("carts.user_id = ?", token).Scan(&inputTransaction)
	if tx.Error != nil {
		log.Fatal("error raw")
		return -1, tx.Error
	}

	data := insert(inputTransaction)

	for i, v := range data {
		data[i].TotalPrice = v.TotalPrice * v.Quantity
		txCreate := repo.db.Create(&data[i])
		if txCreate.Error != nil {
			return -1, txCreate.Error
		}
	}

	return 1, nil

}
