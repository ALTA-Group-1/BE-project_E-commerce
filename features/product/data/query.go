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

func (repo *productData) SelectAllProduct(page int) ([]product.Core, error) {

	var dataProduct []Product
	if page != 0 {
		perPage := 8
		offset := ((page - 1) * perPage)

		queryBuider := repo.db.Limit(perPage).Offset(offset)

		txData := queryBuider.First(&dataProduct)
		return toCoreList(dataProduct), txData.Error

	} else {

		txDataAll := repo.db.Find(&dataProduct)
		return toCoreList(dataProduct), txDataAll.Error

	}

}

func (repo *productData) SelectById(id int) (product.Core, error) {

	var data Product
	tx := repo.db.First(&data, id)

	return data.toCore(), tx.Error

}

func (repo *productData) UpdateData(newData product.Core) (int, error) {

	dataModel := fromCore(newData)

	tx := repo.db.Model(&Product{}).Where("id = ?", newData.ID).Updates(dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil

}

func (repo *productData) DeleteByToken(token int) (int, error) {

	var deleteData User
	tx := repo.db.Delete(&deleteData, token)

	return int(tx.RowsAffected), tx.Error
}

func (repo *productData) SelectMyProduct(token int) ([]product.Core, error) {

	var data []Product
	tx := repo.db.Model(&Product{}).Where("user_id = ?", token).Find(&data)
	return toCoreList(data), tx.Error

}
