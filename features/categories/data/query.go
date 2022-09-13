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
	var dataCart []Results
	tx := repo.db.Model(&Categories{}).Select("products.id, products.images, products.name, products.price").Joins("left join products on products.categories_id = categories.id").Where("products.categories_id = ?", id).Scan(&dataCart)
	if tx.Error != nil {
		return nil, tx.Error
	}

	res := toCoreList(dataCart)
	for key := range res {
		res[key].ID = uint(id)
	}

	return res, nil
}
