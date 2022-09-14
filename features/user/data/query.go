package data

import (
	"errors"
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) InsertData(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectByToken(token int) (user.Core, error) {

	var data User
	tx := repo.db.First(&data, token)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	return data.toCore(), nil

}

func (repo *userData) UpdateData(newData user.Core) (int, error) {
	dataModel := fromCore(newData)

	tx := repo.db.Model(&User{}).Where("id = ?", newData.ID).Updates(dataModel)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("failed to update data")
	}

	return 1, nil
}

func (repo *userData) DeleteByToken(token int) (int, error) {
	var deleteData User
	tx := repo.db.Delete(&deleteData, token)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return int(tx.RowsAffected), nil
}
