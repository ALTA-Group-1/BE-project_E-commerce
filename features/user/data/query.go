package data

import (
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

func (repo *userData) InsertData(data user.Core) (int, error)

func (repo *userData) SelectByToken(token int) (user.Core, error) {

	var data User
	tx := repo.db.First(&data, token)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	userId := data.toCore()
	return userId, nil

}
