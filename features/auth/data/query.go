package data

import (
	"log"
	"project/e-commerce/features/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.DataInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) LoginUser(email string) (auth.Core, error) {

	var data User
	txEmail := repo.db.Where("email = ?", email).First(&data)
	if txEmail.Error != nil {
		log.Println("Error tx")
		return auth.Core{}, txEmail.Error
	}

	if txEmail.RowsAffected != 1 {
		log.Println("Error row")
		return auth.Core{}, txEmail.Error
	}

	var dataUser = toCore(data)

	return dataUser, nil

}
