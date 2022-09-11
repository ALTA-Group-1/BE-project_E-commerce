package data

import (
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
}

func fromCore(dataCore user.Core) User {
	dataModel := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
	}
	return dataModel
}

func (data *User) toCore() user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
	}
}

func toCoreList(data []User) []user.Core {
	var dataCore []user.Core
	for key, _ := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
