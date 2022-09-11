package delivery

import "project/e-commerce/features/auth"

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data AuthRequest) auth.Core {
	return auth.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}
