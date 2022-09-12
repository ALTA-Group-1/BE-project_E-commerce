package delivery

import (
	"project/e-commerce/features/product"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type ProductDelivery struct {
	productUsecase product.UsecaseInterface
}

func New(e *echo.Echo, usecase product.UsecaseInterface) {
	handler := &ProductDelivery{
		productUsecase: usecase,
	}

	e.DELETE("/users", handler.DeleteProduct, middlewares.JWTMiddleware())
}

func (delivery *ProductDelivery) DeleteProduct(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	row, err := delivery.productUsecase.DeleteData(idToken)
	if err != nil || row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))
}
