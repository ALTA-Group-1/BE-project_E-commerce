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
	e.GET("/products", handler.GetAllPagination)
	e.GET("/products/:id", handler.GetProductById)
	e.PUT("/products", handler.PutData, middlewares.JWTMiddleware())
	e.DELETE("/products", handler.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/myproducts", handler.GetAllMyProduct, middlewares.JWTMiddleware())

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

	dataProduct.UserID = idToken

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
	if query == "" {
		query = "0"
	}
	page, err := strconv.Atoi(query)
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("query param must be number"))
	}

	data, errGet := delivery.productUsecase.GetAllProduct(page)
	if errGet != nil {
		return c.JSON(400, helper.FailedResponseHelper("error get all data"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("product data is still empty"))
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

func (delivery *ProductDelivery) PutData(c echo.Context) error {
	var dataUpdate ProductRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind data"))
	}

	var add product.Core
	if dataUpdate.Images != "" {
		add.Images = dataUpdate.Images
	}
	if dataUpdate.Name != "" {
		add.Name = dataUpdate.Name
	}
	if dataUpdate.Stock != 0 {
		add.Stock = dataUpdate.Stock
	}
	if dataUpdate.Price != 0 {
		add.Price = dataUpdate.Price
	}
	if dataUpdate.Desc != "" {
		add.Desc = dataUpdate.Desc
	}
	if dataUpdate.Categories != 0 {
		add.CategoriesID = dataUpdate.Categories
	}

	idToken := middlewares.ExtractToken(c)
	add.ID = uint(idToken)

	row, err := delivery.productUsecase.PutData(add)
	if err != nil || row < 1 {
		return c.JSON(400, helper.FailedResponseHelper("Bad Request"))
	}

	return c.JSON(200, helper.SuccessResponseHelper("Successful Operation"))
}

func (delivery *ProductDelivery) DeleteProduct(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	row, err := delivery.productUsecase.DeleteData(idToken)
	if err != nil || row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("wrong token"))
	}
	return c.JSON(200, helper.SuccessResponseHelper("succes delete"))
}

func (delivery *ProductDelivery) GetAllMyProduct(c echo.Context) error {

	idToken := middlewares.ExtractToken(c)
	data, err := delivery.productUsecase.GetMyProduct(idToken)
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error get my product"))
	} else if len(data) == 0 {
		return c.JSON(200, helper.SuccessResponseHelper("you don't have product data"))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper("success get your product", fromCoreListMyProduct(data)))

}
