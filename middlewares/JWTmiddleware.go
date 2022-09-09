package middlewares

import (
	"project/e-commerce/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(config.SECRET_JWT),
	})

}

func CreateToken(userId int) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.SECRET_JWT))

}

func ExtractToken(c echo.Context) int {

	user := c.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}

	return 0
}
