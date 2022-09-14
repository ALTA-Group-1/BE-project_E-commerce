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
	e.GET("/carts", handler.GetAllCart, middlewares.JWTMiddleware())
	e.PUT("/cart/:id", handler.UpdatePlus, middlewares.JWTMiddleware())
	e.PUT("/cart/:id", handler.UpdateMinus, middlewares.JWTMiddleware())
}

func (delivery *CartDelivery) DeleteCart(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	userID := idToken

	id := c.Param("id")
	idCnv, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	cartID := idCnv

	row, err := delivery.cartUsecase.DeleteCart(userID, cartID)
	if err != nil || row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))
}

func (delivery *CartDelivery) GetAllCart(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)
	data, err := delivery.cartUsecase.GetByToken(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get cart"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("you dont have product in cart"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("succes get cart", data))

}

func (delivery *CartDelivery) UpdatePlus(c echo.Context) error {
	id := c.Param("id")
	idCnv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be a number"))
	}
	errBind := c.Bind(&idCnv)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	cartID := idCnv

	query := c.QueryParam("increment")
	if query == "" {
		query = "0"
	}

	row, err := delivery.cartUsecase.UpdatePlus(cartID, query)
	if err != nil || row < 1 {
		return c.JSON(400, helper.FailedResponseHelper("Bad Request"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Successful Operation"))
}

func (delivery *CartDelivery) UpdateMinus(c echo.Context) error {
	id := c.Param("id")
	idCnv, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be a number"))
	}
	errBind := c.Bind(&idCnv)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	cartID := idCnv

	query := c.QueryParam("decrement")
	if query == "" {
		query = "0"
	}

	row, err := delivery.cartUsecase.UpdateMinus(cartID, query)
	if err != nil || row < 1 {
		return c.JSON(400, helper.FailedResponseHelper("Bad Request"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Successful Operation"))
}
