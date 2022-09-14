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
	// post cart
	// get cart
	GetByToken(token int) ([]Core, error)
	UpdatePlus(cartID int, increment string) (int, error)
	UpdateMinus(cartID int, decrement string) (int, error)
	DeleteCart(userID, cartID int) (int, error)
}

type DataInterface interface {
	// post cart
	// get cart
	SelectByToken(token int) ([]Core, error)
	UpdatePlusData(cartID int, increment string) (int, error)
	UpdateMinusData(cartID int, decrement string) (int, error)
	DeleteData(userID, cartID int) (int, error)
}
