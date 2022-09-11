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

// func (repo *userData) InsertData(data user.Core) (int, error)

func (repo *userData) SelectByToken(token int) (user.Core, error) {

	var data User
	tx := repo.db.First(&data, token)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	userId := data.toCore()
	return userId, nil

}

func (repo *userData) UpdateData(id int, newData user.Core) (int, error) {
	dataModel := fromCore(newData)

	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil
}
