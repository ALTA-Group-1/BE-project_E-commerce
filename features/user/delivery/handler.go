package delivery

import (
	"project/e-commerce/features/user"

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
