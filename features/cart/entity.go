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
	ID     uint
	Name   string
	Images string
	Price  int
}

type UsecaseInterface interface {
	PostData(data Core) (int, error)
	// get cart
	GetByToken(token int) ([]Core, error)
	// update cart
	DeleteCart(userID, cartID int) (int, error)
}

type DataInterface interface {
	InsertData(data Core) (int, error)
	// get cart
	SelectByToken(token int) ([]Core, error)
	// update cart
	DeleteData(userID, cartID int) (int, error)
}
