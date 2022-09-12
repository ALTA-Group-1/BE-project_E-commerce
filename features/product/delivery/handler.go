package delivery

import (
	"project/e-commerce/features/product"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductDelivery struct {
	productUsecase product.UsecaseInterface
}

func New(e *echo.Echo, usecase product.UsecaseInterface) {
	handler := &ProductDelivery{
		productUsecase: usecase,
	}

	e.POST("/products", handler.PostData, middlewares.JWTMiddleware())
	e.DELETE("/products", handler.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/products", handler.GetAllPagination)
	e.GET("/products/:id", handler.GetProductById)

}

func (delivery *ProductDelivery) DeleteProduct(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	row, err := delivery.productUsecase.DeleteData(idToken)
	if err != nil || row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))

}

func (delivery *ProductDelivery) PostData(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == 0 {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}

	var dataProduct ProductRequest
	errBind := c.Bind(&dataProduct)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.productUsecase.PostData(toCore(dataProduct))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(400, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *ProductDelivery) GetAllPagination(c echo.Context) error {

	query := c.QueryParam("page")
	page, err := strconv.Atoi(query)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	}

	data, errGet := delivery.productUsecase.GetAllProduct(page)
	if errGet != nil || len(data) == 0 {
		return c.JSON(400, helper.FailedResponseHelper("error get all data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get all data", fromCoreList(data)))

}

func (delivery *ProductDelivery) GetProductById(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return c.JSON(400, helper.FailedResponseHelper("param must be number > 0"))
	}

	data, errFind := delivery.productUsecase.GetById(id)
	if errFind != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get by id"))
	} else if data.Name == "" {
		return c.JSON(400, helper.FailedResponseHelper("data not found"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("succes get by id", fromCore(data)))

}
