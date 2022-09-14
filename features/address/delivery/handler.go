package delivery

import (
	"project/e-commerce/features/address"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type addressDelivery struct {
	addressUsecase address.UsecaseInterface
}

func New(e *echo.Echo, usecase address.UsecaseInterface) {
	handler := &addressDelivery{
		addressUsecase: usecase,
	}
	e.POST("/address", handler.PostDataAddress, middlewares.JWTMiddleware())

}

func (delivery *addressDelivery) PostDataAddress(c echo.Context) error {
	var dataReq Request
	err := c.Bind(&dataReq)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error request"))
	}

	idToken := middlewares.ExtractToken(c)

	request := dataReq.fromCore()
	_, errAdd := delivery.addressUsecase.PostData(idToken, request)
	if errAdd != nil {
		return c.JSON(400, helper.FailedResponseHelper("error orders"))
	} else {
		return c.JSON(200, helper.SuccessResponseHelper("succes orders"))
	}

}
