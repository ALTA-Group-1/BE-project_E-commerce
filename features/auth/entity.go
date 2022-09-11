package auth

type Core struct {
	ID       uint
	Email    string
	Password string
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) string
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
}
