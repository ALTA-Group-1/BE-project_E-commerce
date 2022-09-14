package data

import (
	"project/e-commerce/features/transaction"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Quantity    uint
	TotalPrice  uint
	OrderStatus string
	CartID      uint
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

type result struct {
	Quantity   uint
	TotalPrice uint
	CartID     uint
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
