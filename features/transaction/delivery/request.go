package delivery

import "project/e-commerce/features/transaction"

type Request struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	PostCode uint   `json:"postcode" form:"postcode"`
	Visa     string `json:"visa" form:"visa"`
	Name     string `json:"name on card" form:"name on card"`
	Number   uint   `json:"number card" form:"number card"`
	Cvv2     uint   `json:"noCvv2" form:"noCvv2"`
	Month    uint   `json:"month" form:"month"`
	Year     uint   `json:"year" form:"year"`
}

func (data *Request) fromCore() (transaction.AddressCore, transaction.PaymentCore) {
	dataAddress := transaction.AddressCore{
		Street:   data.Street,
		City:     data.City,
		Province: data.Province,
		PostCode: data.PostCode,
	}

	dataPayment := transaction.PaymentCore{
		Visa:   data.Visa,
		Name:   data.Name,
		Number: data.Number,
		Cvv2:   data.Cvv2,
		Month:  data.Month,
		Year:   data.Year,
	}

	return dataAddress, dataPayment

}
