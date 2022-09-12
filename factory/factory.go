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
	// categoriesData "project/e-commerce/features/categories/data"
	// categoriesDelivery "project/e-commerce/features/categories/delivery"
	// categoriesUsecase "project/e-commerce/features/categories/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	// categoriesDataFactory := categoriesData.New(db)
	// categoriesUsecaseFactory := categoriesUsecase.New(categoriesDataFactory)
	// categoriesDelivery.New(e, categoriesUsecaseFactory)
}
