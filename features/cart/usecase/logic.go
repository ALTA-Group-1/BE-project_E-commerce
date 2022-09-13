package usecase

import (
	"errors"
	"project/e-commerce/features/cart"
)

type cartUsecase struct {
	cartData cart.DataInterface
}

func New(data cart.DataInterface) cart.UsecaseInterface {
	return &cartUsecase{
		cartData: data,
	}
}

func (usecase *cartUsecase) PostData(data cart.Core) (int, error) {
	if data.ProductID == 0 {
		return -1, errors.New("data tidak boleh kosong")
	}

	row, err := usecase.cartData.InsertData(data)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *cartUsecase) DeleteCart(userID, cartID int) (int, error) {
	row, err := usecase.cartData.DeleteData(userID, cartID)
	if err != nil {
		return -1, err
	}
	return row, nil
}

func (usecase *cartUsecase) GetByToken(token int) ([]cart.Core, error) {
	res, err := usecase.cartData.SelectByToken(token)
	if err != nil {
		return nil, err
	}

	return res, nil
}
