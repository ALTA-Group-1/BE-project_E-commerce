package delivery

import (
	"project/e-commerce/features/transaction"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type TransactionDelivery struct {
	transactionUsecase transaction.UsecaseInterface
}

func New(e *echo.Echo, usecase transaction.UsecaseInterface) {
	handler := &TransactionDelivery{
		transactionUsecase: usecase,
	}

	e.POST("/orders", handler.PostDataOrders, middlewares.JWTMiddleware())
	e.PUT("/orders/confirm", handler.PostDataOrders, middlewares.JWTMiddleware())

}

func (delivery *TransactionDelivery) PostDataOrders(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)
	var data Request
	errBind := c.Bind(&data)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error request"))
	}

	requestAddress, requestPayment := data.fromCore()

	row, err := delivery.transactionUsecase.PostData(idtoken, requestAddress, requestPayment)
	if err != nil || row == 0 {
		return c.JSON(500, helper.FailedResponseHelper("failed post order"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("succes post data status waiting"))

}

func (delivery *TransactionDelivery) PutStatusConfirm(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)

	row, err := delivery.transactionUsecase.PutStatus(idtoken, "confirm")
	if err != nil || row == 0 {
		return c.JSON(400, helper.FailedResponseHelper("low stock product"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("succes orders products"))

}
