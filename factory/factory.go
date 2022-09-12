package factory

import (
	userData "project/e-commerce/features/user/data"
	userDelivery "project/e-commerce/features/user/delivery"
	userUsecase "project/e-commerce/features/user/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authData "project/e-commerce/features/auth/data"
	authDelivery "project/e-commerce/features/auth/delivery"
	authUsecase "project/e-commerce/features/auth/usecase"

	productData "project/e-commerce/features/product/data"
	productDelivery "project/e-commerce/features/product/delivery"
	productUsecase "project/e-commerce/features/product/usecase"
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
}
