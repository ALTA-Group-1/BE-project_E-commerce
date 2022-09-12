package data

import (
	"errors"
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

func (repo *productData) InsertData(data product.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)

	return int(tx.RowsAffected), tx.Error
}

func (repo *productData) DeleteByToken(token int) (int, error) {
	var deleteData User
	tx := repo.db.Delete(&deleteData, token)

	return int(tx.RowsAffected), tx.Error
}

func (repo *productData) SelectAllProduct(page int) ([]product.Core, error) {

	var maksOffset int
	tx := repo.db.Raw("SELECT COUNT * FROM products WHERE deleted_at = NULL").Scan(&maksOffset)
	if tx.Error != nil {
		return nil, tx.Error
	}

	perPage := 8
	offset := ((page - 1) * perPage)

	if offset > maksOffset {
		return nil, errors.New("maks page")
	}

	var dataProduct []Product
	txData := repo.db.Raw("SELECT * FROM products WHERE deleted_at = NULL ORDER BY name ASC LIMIT = ? OFFSET = ?", perPage, offset).Scan(&dataProduct)

	return toCoreList(dataProduct), txData.Error

}

func (repo *productData) SelectById(id int) (product.Core, error) {

	var data Product
	tx := repo.db.First(&data, id)

	return data.toCore(), tx.Error

}
