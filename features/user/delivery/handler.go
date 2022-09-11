package delivery

import (
	"project/e-commerce/features/user"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.POST("/users", handler.PostData)

}

func (delivery *UserDelivery) PostData(c echo.Context) error

func (users *UserDelivery) GetByTokenJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)

	res, err := users.userUsecase.GetByToken(idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("error token"))
	}

	respon := fromCore(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}
