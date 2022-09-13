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
	// get all product
	GetAllProduct(page int) ([]Core, error)
	// get detail product
	// update prdocut
	GetById(param int) (Core, error)
	// update product
	DeleteData(token int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (row int, err error)
	// get all product
	SelectAllProduct(page int) ([]Core, error)
	// get detail product
	SelectById(param int) (Core, error)
	// update product
	DeleteByToken(token int) (int, error)
}
