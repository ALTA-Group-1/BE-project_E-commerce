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

	e.POST("/transactions", handler.PostData, middlewares.JWTMiddleware())

}

func (deliver *TransactionDelivery) PostData(c echo.Context) error {

	idtoken := middlewares.ExtractToken(c)

	row, err := deliver.transactionUsecase.PostData(idtoken)
	if err != nil || row == 0 {
		return c.JSON(500, helper.FailedResponseHelper("failed post order"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("succes post data status waiting"))

}
