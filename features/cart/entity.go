package cart

type Core struct {
	ID            uint
	Quantity      int
	ProductID     uint
	ProductName   string
	ProductImages string
	ProductPrice  int
	UserID        int
}

type Product struct {
	ID       uint
	Name     string
	Images   string
	Price    int
	Quantity int
}

type UsecaseInterface interface {
	PostData(data Core) (int, error)
	GetByToken(token int) ([]Core, error)
	PutData(cartID, token int, update string) (int, error)
	DeleteCart(userID, cartID int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (int, error)
	SelectByToken(token int) ([]Core, error)
	UpdateData(cartID, token int, update string) (int, error)
	DeleteData(userID, cartID int) (int, error)
}
