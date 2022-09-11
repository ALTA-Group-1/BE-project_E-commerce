package data

import (
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
