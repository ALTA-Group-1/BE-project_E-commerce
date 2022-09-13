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
	e.POST("/carts", handler.PostData, middlewares.JWTMiddleware())
	e.GET("/carts", handler.GetAllCart, middlewares.JWTMiddleware())
	e.DELETE("/carts/:id", handler.DeleteCart, middlewares.JWTMiddleware())
}

func (delivery *CartDelivery) PostData(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == 0 {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}

	var dataCart CartRequest
	errBind := c.Bind(&dataCart)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	dataCart.Quantity = 1
	dataCart.UserID = idToken

	row, err := delivery.cartUsecase.PostData(toCore(dataCart))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
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

	dataCart, totalRes := fromCoreList(data)
	return c.JSON(200, helper.SuccessDataResponseHelper(totalRes, dataCart))

}
