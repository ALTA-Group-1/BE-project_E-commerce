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

type PaymentCore struct {
	TransactionID uint
	Visa          string
	Name          string
	Number        uint
	Cvv2          uint
	Month         uint
	Year          uint
}

type UsecaseInterface interface {
	PostData(token int, data AddressCore, dataPay PaymentCore) (int, error)
	PutStatus(token int, status string) (int, error)
	DeleteOrder(token int, status string) (int, error)
}

type DataInterface interface {
	InsertData(token int, data AddressCore, dataPay PaymentCore) (int, error)
	UpdateStatus(token int, status string) (int, error)
	CancelOrder(token int, status string) (int, error)
}
