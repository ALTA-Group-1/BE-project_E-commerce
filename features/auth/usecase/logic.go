package usecase

import (
	"project/e-commerce/features/auth"
	"project/e-commerce/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	authData auth.DataInterface
}

func New(data auth.DataInterface) auth.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) LoginAuthorized(email, password string) string {

	if email == "" || password == "" {
		return "please input email and password"
	}

	results, errEmail := usecase.authData.LoginUser(email)
	if errEmail != nil {
		return "email not found"
	}

	errPw := bcrypt.CompareHashAndPassword([]byte(results.Password), []byte(password))
	if errPw != nil {
		return "wrong password"
	}

	token, errToken := middlewares.CreateToken(int(results.ID))

	if errToken != nil {
		return "error to created token"
	}

	return token

}
