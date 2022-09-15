package data

import (
	"project/e-commerce/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Quantity    int
	TotalPrice  int
	OrderStatus string
	CartID      uint
}

type Address struct {
	TransactionID uint `gorm:"primary_key;ForeignKey:TransactionID"`
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type Product struct {
	gorm.Model
	Name         string
	Images       string
	Price        int
	Stock        int
	Desc         string
	UserID       uint
	CategoriesID uint
	CartID       []Cart `gorm:"foreignKey:ProductID"`
}

type Cart struct {
	gorm.Model
	Quantity  int
	ProductID uint
	UserID    uint
}

type Results struct {
	ID        uint
	Quantity  int
	Name      string
	Images    string
	Price     int
	UserID    uint
	ProductID uint
}

type result struct {
	CartID     uint
	Quantity   int
	TotalPrice int
}

type DBTransaction struct {
	Cart_id  uint
	Quantity int
	Stock    int
	ID       uint
}

func insertJoin(data []Results) []result {
	var dataRes []result
	for _, v := range data {
		dataRes = append(dataRes, result{
			Quantity:   v.Quantity,
			TotalPrice: v.Price,
			CartID:     v.ID,
		})
	}

	return dataRes
}

func insert(res []result) []Transaction {
	var data []Transaction
	for _, v := range res {
		data = append(data, Transaction{
			Quantity:    v.Quantity,
			TotalPrice:  v.TotalPrice,
			CartID:      v.CartID,
			OrderStatus: "waiting",
		})
	}

	return data
}

func (tx *Transaction) toCore() transaction.Core {
	return transaction.Core{
		ID:          tx.ID,
		Quantity:    tx.Quantity,
		TotalPrice:  tx.TotalPrice,
		OrderStatus: tx.OrderStatus,
		CartID:      tx.CartID,
		CreatedAt:   tx.CreatedAt,
		Updated:     tx.UpdatedAt,
		DeletedAt:   tx.DeletedAt.Time,
	}
}

func toDb(data transaction.AddressCore) Address {
	return Address{
		TransactionID: data.TransactionID,
		Street:        data.Street,
		City:          data.City,
		Province:      data.Province,
		PostCode:      data.PostCode,
	}

}
