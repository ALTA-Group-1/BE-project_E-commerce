package delivery

import (
	"net/http"
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

	// e.POST("/users", handler.PostData)
	e.GET("/users", handler.GetByTokenJWT, middlewares.JWTMiddleware())
	e.PUT("/users", handler.PutData, middlewares.JWTMiddleware())
}

// func (delivery *UserDelivery) PostData(c echo.Context) error

func (users *UserDelivery) GetByTokenJWT(e echo.Context) error {

	idToken := middlewares.ExtractToken(e)

	res, err := users.userUsecase.GetByToken(idToken)
	if err != nil {
		return e.JSON(400, helper.FailedResponseHelper("error token"))
	}

	respon := fromCore(res)

	return e.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))

}

func (delivery *UserDelivery) PutData(c echo.Context) error {
	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error bind data"))
	}

	idToken := middlewares.ExtractToken(c)

	row, err := delivery.userUsecase.PutData(idToken, toCore(dataUpdate))
	if err != nil || row < 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("Bad Request"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("Successful Operation"))
}
