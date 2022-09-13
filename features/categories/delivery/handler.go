package delivery

import (
	"project/e-commerce/features/categories"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoriesDelivery struct {
	categoriesUsecase categories.UsecaseInterface
}

func New(e *echo.Echo, usecase categories.UsecaseInterface) {
	handler := &CategoriesDelivery{
		categoriesUsecase: usecase,
	}
	e.GET("/categories/:id", handler.GetAll)
}

func (delivery *CategoriesDelivery) GetAll(c echo.Context) error {
	id := c.Param("id")
	categoryId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("param must be number"))
	}

	results, err := delivery.categoriesUsecase.GetAll(categoryId)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success", fromCoreList(results)))
}
