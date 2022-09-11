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
	e.GET("/users", handler.GetByTokenJWT, middlewares.JWTMiddleware())
	e.PUT("/users", handler.PutData, middlewares.JWTMiddleware())
}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRegister UserRequest
	errBind := c.Bind(&dataRegister)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}
	row, err := delivery.userUsecase.PostData(toCore(dataRegister))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetByTokenJWT(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	res, err := delivery.userUsecase.GetByToken(idToken)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error token"))
	}

	respon := fromCore(res)

	return c.JSON(200, helper.SuccessDataResponseHelper("succes get data by id", respon))
}

func (delivery *UserDelivery) PutData(c echo.Context) error {
	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	var add user.Core
	if dataUpdate.Email != "" {
		add.Email = dataUpdate.Email
	}
	if dataUpdate.Name != "" {
		add.Name = dataUpdate.Name
	}
	if dataUpdate.Password != "" {
		add.Password = dataUpdate.Password
	}
	if dataUpdate.Phone != "" {
		add.Phone = dataUpdate.Phone
	}
	if dataUpdate.Address != "" {
		add.Address = dataUpdate.Address
	}

	idToken := middlewares.ExtractToken(c)
	add.ID = uint(idToken)

	row, err := delivery.userUsecase.PutData(add)
	if err != nil || row < 1 {
		return c.JSON(400, helper.FailedResponseHelper("Bad Request"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Successful Operation"))
}

func (delivery *UserDelivery) DeleteUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	row, err := delivery.userUsecase.DeleteData(idToken)
	if err != nil || row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))
}
