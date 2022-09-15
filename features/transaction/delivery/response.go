package delivery

import "project/e-commerce/features/transaction"

type ResponHistoryOrder struct {
	Images     string
	Name       string
	Price      uint
	Quantity   uint
	TotalPrice uint
}

func toRespon(data []transaction.HistoryOrder) []ResponHistoryOrder {
	var dataRes []ResponHistoryOrder
	for i, v := range data {
		dataRes = append(dataRes, ResponHistoryOrder{
			Images:   v.Images,
			Name:     v.Name,
			Price:    v.Price,
			Quantity: v.Quantity,
		})

		dataRes[i].TotalPrice = v.Price * v.Quantity

	}

	return dataRes
}
