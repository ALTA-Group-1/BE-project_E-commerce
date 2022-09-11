package usecase

import (
	"project/e-commerce/features/auth"
)

type authUsecase struct {
	authData auth.DataInterface
}

func New(data auth.DataInterface) auth.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}
