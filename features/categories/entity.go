package categories

type Core struct {
	ID            uint
	Name          string
	ProductID     uint
	ProductName   string
	ProductImages string
	ProductPrice  int
	ProductStock  int
	ProductDesc   string
}

type ProductCore struct {
	ID     uint
	Name   string
	Images string
	Price  int
	Stock  int
	Desc   string
}

type UsecaseInterface interface {
	GetAll(id int) (data []Core, err error)
}

type DataInterface interface {
	GetAllData(id int) (data []Core, err error)
}
