package delivery

import "project/e-commerce/features/transaction"

type Request struct {
	Street   string `json:"street" form:"street"`
	City     string `json:"city" form:"city"`
	Province string `json:"province" form:"province"`
	PostCode uint   `json:"postcode" form:"postcode"`
}

func (data *Request) fromCore() transaction.AddressCore {
	return transaction.AddressCore{
		Street:   data.Street,
		City:     data.City,
		Province: data.Province,
		PostCode: data.PostCode,
	}

}
