package categories

import "project/e-commerce/features/product"

type Core struct {
	ID      uint
	Name    string
	Product product.Core
}
