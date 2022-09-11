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
	// create user
	PostData(data Core) (int, error)
	// read user
	// read user by id
	GetByToken(token int) (data Core, err error)
	// put user
	// delete user
}

type DataInterface interface {
	// create user
	InsertData(data Core) (int, error)
	// read user by id
	SelectByToken(token int) (data Core, err error)
	// put user
	// delete user
}
