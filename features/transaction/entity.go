package transaction

import (
	"time"
)

type Core struct {
	ID          uint
	Quantity    int
	TotalPrice  int
	OrderStatus string
	CartID      uint
	CreatedAt   time.Time
	Updated     time.Time
	DeletedAt   time.Time
}

type AddressCore struct {
	TransactionID uint
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type UsecaseInterface interface {
	PostData(token int, data AddressCore) (int, error)
	PutStatus(token int, status string) (int, error)
  // DeleteOrder()
}

type DataInterface interface {
	InsertData(token int, data AddressCore) (int, error)
	UpdateStatus(token int, status string) (int, error)
	// DeleteData()
}
