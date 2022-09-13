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
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *productData) SelectAllProduct(page int) ([]product.Core, error) {

	var dataProduct []Product
	if page != 0 {
		perPage := 8
		offset := ((page - 1) * perPage)

		queryBuider := repo.db.Limit(perPage).Offset(offset)

		txData := queryBuider.First(&dataProduct)
		if txData.Error != nil {
			return nil, txData.Error
		}
		return toCoreList(dataProduct), nil

	} else {

		txDataAll := repo.db.Find(&dataProduct)
		if txDataAll.Error != nil {
			return nil, txDataAll.Error
		}
		return toCoreList(dataProduct), nil
	}

}

func (repo *productData) SelectById(id int) (product.Core, error) {
	var data Product
	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return product.Core{}, tx.Error
	}
	return data.toCore(), nil
}

func (repo *productData) UpdateData(token int, newData product.Core) (int, error) {
	dataModel := fromCore(newData)

	tx := repo.db.Model(&Product{}).Where("id = ? AND user_id = ?", newData.ID, token).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("failed to update data")
	}

	return 1, nil
}

func (repo *productData) DeleteByToken(param, token int) (int, error) {
	var deleteData Product
	tx := repo.db.First(&deleteData, param)
	if tx.Error != nil {
		return -1, tx.Error
	}

	productId := deleteData.toCore()

	if productId.UserID == token {
		var data Product
		txDelId := repo.db.Delete(&data, param)
		if txDelId.Error != nil {
			return -1, txDelId.Error
		}
		var err error
		return int(txDelId.RowsAffected), err
	} else {
		return -1, errors.New("no access")
	}
	// tx := repo.db.Delete(&deleteData, token)

	// return int(tx.RowsAffected), tx.Error
}

func (repo *productData) SelectMyProduct(token int) ([]product.Core, error) {
	var data []Product
	tx := repo.db.Model(&Product{}).Where("user_id = ?", token).Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return toCoreList(data), nil
}
