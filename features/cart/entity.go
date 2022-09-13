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
	// post cart
	// get cart
	// update cart
	// delete cart
}

type DataInterface interface {
	// post cart
	// get cart
	// update cart
	// delete cart
}
