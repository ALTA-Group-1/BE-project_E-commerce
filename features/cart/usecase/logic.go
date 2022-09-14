package usecase

import (
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

func (usecase *cartUsecase) DeleteCart(userID, cartID int) (int, error) {
	row, err := usecase.cartData.DeleteData(userID, cartID)
	if err != nil {
		return 0, err
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

func (usecase *cartUsecase) UpdatePlus(cartID int, increment string) (int, error) {
	row, err := usecase.cartData.UpdatePlusData(cartID, increment)
	if err != nil {
		return 0, err
	}
	return row, nil
}

func (usecase *cartUsecase) UpdateMinus(cartID int, decrement string) (int, error) {
	row, err := usecase.cartData.UpdateMinusData(cartID, decrement)
	if err != nil {
		return 0, err
	}
	return row, nil
}
