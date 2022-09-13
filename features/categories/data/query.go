package data

import (
	"project/e-commerce/features/categories"

	"gorm.io/gorm"
)

type categoriesData struct {
	db *gorm.DB
}

func New(db *gorm.DB) categories.DataInterface {
	return &categoriesData{
		db: db,
	}
}

func (repo *categoriesData) GetAllData(id int) ([]categories.Core, error) {
	limit := 8
	offset := ((id - 1) * limit)

	queryBuider := repo.db.Limit(limit).Offset(offset)

	var dataCategories []Categories
	// tx := queryBuider.Model(&Product{}).Where("id = ?", id).Find(&dataCategories)
	tx := queryBuider.Where("id = ?", id).Preload("Product").Find(&dataCategories)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(dataCategories)
	return dataCore, nil
}
