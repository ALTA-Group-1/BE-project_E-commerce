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

func (usecase *transactionUsecase) PostData(token int, data transaction.AddressCore) (int, error) {

	row, err := usecase.transactionData.InsertData(token, data)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *transactionUsecase) PutStatus(token int, status string) (int, error) {

	row, err := usecase.transactionData.UpdateStatus(token, status)
	if err != nil {
		return -1, err
	}

	return row, nil

}
