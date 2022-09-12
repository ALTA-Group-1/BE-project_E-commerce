package product

import (
	"project/e-commerce/features/user"
	"time"
)

type Core struct {
	ID        uint
	Name      string
	Images    string
	Price     int
	Stock     int
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
}

type UsecaseInterface interface {
	// create product
	// get all product
	// get detail product
	// get category
	// update product
	// delete product
}

type DataInterface interface {
	// create product
	// get all product
	// get detail product
	// update product
	// delete product
}