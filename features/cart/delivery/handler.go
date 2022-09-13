package delivery

import (
	"project/e-commerce/features/cart"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartDelivery struct {
	cartUsecase cart.UsecaseInterface
}

func New(e *echo.Echo, usecase cart.UsecaseInterface) {
	handler := &CartDelivery{
		cartUsecase: usecase,
	}
	e.DELETE("/cart/:id", handler.DeleteCart, middlewares.JWTMiddleware())
}

func (delivery *CartDelivery) DeleteCart(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	userID := idToken

	id := c.Param("productID")
	idCnv, _ := strconv.Atoi(id)
	productID := idCnv

	row, err := delivery.cartUsecase.DeleteCart(userID, productID)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))
}
