package usecase

import (
	"errors"
	"project/e-commerce/features/product"
)

type productUsecase struct {
	productData product.DataInterface
}

func New(data product.DataInterface) product.UsecaseInterface {
	return &productUsecase{
		productData: data,
	}
}

func (usecase *productUsecase) PostData(data product.Core) (int, error) {
	if data.Name == "" || data.Images == "" || data.Price == 0 || data.Stock == 0 || data.Desc == "" || data.CategoriesID == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.productData.InsertData(data)
	return row, err

}

func (usecase *productUsecase) GetAllProduct(page int) ([]product.Core, error) {

	data, err := usecase.productData.SelectAllProduct(page)
	return data, err

}

func (usecase *productUsecase) GetById(id int) (product.Core, error) {

	data, err := usecase.productData.SelectById(id)
	return data, err

}

func (usecase *productUsecase) PutData(newData product.Core) (int, error) {
	row, err := usecase.productData.UpdateData(newData)

	return row, err
}

func (usecase *productUsecase) DeleteData(token int) (int, error) {
	row, err := usecase.productData.DeleteByToken(token)

	return row, err
}

func (usecase *productUsecase) GetMyProduct(token int) ([]product.Core, error) {

	data, err := usecase.GetMyProduct(token)
	return data, err

}
