package user

import (
	"time"
)

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductCore struct {
	ID        uint
	Name      string
	Images    string
	Price     string
	Stock     int
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UsecaseInterface interface {
	PostData(data Core) (int, error)
	GetByToken(token int) (data Core, err error)
	PutData(newData Core) (row int, err error)
	DeleteData(token int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (int, error)
	SelectByToken(token int) (data Core, err error)
	UpdateData(newData Core) (row int, err error)
	DeleteByToken(token int) (int, error)
}
