package data

import (
	"project/e-commerce/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(user User) auth.Core {

	var core = auth.Core{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	return core
}
