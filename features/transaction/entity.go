package transaction

import (
	"time"
)

type Core struct {
	ID          uint
	Quantity    uint
	TotalPrice  uint
	OrderStatus string
	CartID      uint
	CreatedAt   time.Time
	Updated     time.Time
	DeletedAt   time.Time
}

type UsecaseInterface interface {
	PostData(token int) (int, error)
}

type DataInterface interface {
	InsertData(token int) (int, error)
}
