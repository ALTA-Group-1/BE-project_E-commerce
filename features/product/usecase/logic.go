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
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *productUsecase) GetAllProduct(page int) ([]product.Core, error) {
	data, err := usecase.productData.SelectAllProduct(page)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (usecase *productUsecase) GetById(id int) (product.Core, error) {
	data, err := usecase.productData.SelectById(id)
	if err != nil {
		return product.Core{}, err
	}
	return data, nil
}

func (usecase *productUsecase) PutData(token int, newData product.Core) (int, error) {
	row, err := usecase.productData.UpdateData(token, newData)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *productUsecase) DeleteData(param, token int) (int, error) {
	row, err := usecase.productData.DeleteByToken(param, token)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *productUsecase) GetMyProduct(token int) ([]product.Core, error) {
	data, err := usecase.productData.SelectMyProduct(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}
