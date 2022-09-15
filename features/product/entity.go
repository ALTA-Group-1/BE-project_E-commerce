package product

import (
	"time"
)

type Core struct {
	ID           uint
	Name         string
	Images       string
	Price        int
	Stock        int
	Desc         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	UserID       int
	CategoriesID int
}

type User struct {
	ID       uint
	UserName string
}

type CategoriesCore struct {
	ID   uint
	Name string
}

type UsecaseInterface interface {
	PostData(data Core) (row int, err error)
	GetAllProduct(page int, category string) ([]Core, error)
	GetMyProduct(token int) ([]Core, error)
	GetById(param int) (Core, error)
	PutData(token int, newData Core) (row int, err error)
	DeleteData(param, token int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	SelectAllProduct(page int, category string) ([]Core, error)
	SelectMyProduct(token int) ([]Core, error)
	SelectById(param int) (Core, error)
	UpdateData(token int, newData Core) (row int, err error)
	DeleteByToken(param, token int) (int, error)
}
