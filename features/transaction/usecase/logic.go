package usecase

import (
	"errors"
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

func (usecase *transactionUsecase) PostData(token int, data transaction.AddressCore, dataPay transaction.PaymentCore) (int, error) {

	if data.City == "" || data.PostCode == 0 || data.Province == "" || data.Street == "" || dataPay.Name == "" || dataPay.Cvv2 == 0 || dataPay.Month == 0 || dataPay.Month > 12 || dataPay.Year == 0 || dataPay.Number == 0 || dataPay.Visa == "" {
		return -1, errors.New("error")
	}

	row, err := usecase.transactionData.InsertData(token, data, dataPay)
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

func (usecase *transactionUsecase) DeleteOrder(token int, status string) (int, error) {
	row, err := usecase.transactionData.CancelOrder(token, status)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func (usecase *transactionUsecase) GetOrder(token int) ([]transaction.HistoryOrder, error) {

	data, err := usecase.transactionData.SelectOrder(token)
	if err != nil {
		return nil, err
	}

	return data, nil
}
