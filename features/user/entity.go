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
	// put user
	// delete user
}

type DataInterface interface {
	// create user
	InsertData(data Core) (int, error)
	// read user
	// read user by id
	// put user
	// delete user
}
