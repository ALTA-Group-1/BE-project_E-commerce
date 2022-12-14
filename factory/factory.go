package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "project/e-commerce/features/user/data"
	userDelivery "project/e-commerce/features/user/delivery"
	userUsecase "project/e-commerce/features/user/usecase"

	authData "project/e-commerce/features/auth/data"
	authDelivery "project/e-commerce/features/auth/delivery"
	authUsecase "project/e-commerce/features/auth/usecase"

	productData "project/e-commerce/features/product/data"
	productDelivery "project/e-commerce/features/product/delivery"
	productUsecase "project/e-commerce/features/product/usecase"

	cartData "project/e-commerce/features/cart/data"
	cartDelivery "project/e-commerce/features/cart/delivery"
	cartUsecase "project/e-commerce/features/cart/usecase"

	transactionData "project/e-commerce/features/transaction/data"
	transactionDelivery "project/e-commerce/features/transaction/delivery"
	transactionUsecase "project/e-commerce/features/transaction/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	productDataFactory := productData.New(db)
	productUsecaseFactory := productUsecase.New(productDataFactory)
	productDelivery.New(e, productUsecaseFactory)

	cartDataFactory := cartData.New(db)
	cartUsecaseFactory := cartUsecase.New(cartDataFactory)
	cartDelivery.New(e, cartUsecaseFactory)

	transactionDataFactory := transactionData.New(db)
	transactionUsecaseFactory := transactionUsecase.New(transactionDataFactory)
	transactionDelivery.New(e, transactionUsecaseFactory)

}
