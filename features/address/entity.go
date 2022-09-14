package address

type Core struct {
	TransactionID uint
	Street        string
	City          string
	Province      string
	PostCode      uint
}

type UsecaseInterface interface {
	PostData(token int, data Core) (int, error)
}

type DataInterface interface {
	InsertData(token int, data Core) (int, error)
}
