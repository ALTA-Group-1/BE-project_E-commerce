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

type UsecaseInterface interface {
	PostData(data Core) (int, error)
	GetByToken(token int) (data Core, err error)
	PutData(newData Core) (row int, err error)
	DeleteData(token int) (int, error)
}

type DataInterface interface {
	// create user
	InsertData(data Core) (int, error)
	// read user by id
	SelectByToken(token int) (data Core, err error)
	// put user
	UpdateData(newData Core) (row int, err error)
	// delete user
	DeleteByToken(token int) (int, error)
}
