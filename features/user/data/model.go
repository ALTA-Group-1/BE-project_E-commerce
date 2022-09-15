package data

import (
	"project/e-commerce/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Phone    string
	Address  string
	Product  []Product `gorm:"foreignKey:UserID"`
	Cart     []Cart    `gorm:"foreignKey:UserID"`
}

type Product struct {
	gorm.Model
	Name         string
	Images       string
	Price        int
	Stock        int
	Desc         string
	UserID       uint
	CategoriesID uint
}

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	UserID    uint
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
