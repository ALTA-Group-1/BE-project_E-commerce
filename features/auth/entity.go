package auth

type Core struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UsecaseInterface interface {
}

type DataInterface interface {
}
