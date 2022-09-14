package usecase

import (
	"project/e-commerce/features/transaction"
)

type transactionUsecase struct {
	transactionData transaction.DataInterface
}

func New(data transaction.DataInterface) transaction.UsecaseInterface {
	return &transactionUsecase{
		transactionData: data,
	}
}

func (usecase *transactionUsecase) PostData(token int) (int, error) {

	row, err := usecase.transactionData.InsertData(token)
	if err != nil {
		return -1, err
	}

	return row, nil
}
