package middlewares

import (
	"time"

	"github.com/dakasakti/deploy-apps-hexagonal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateToken(id uint, username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetConfig().Secret_JWT))
}

func ExtractToken(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)

	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		id := claims["user_id"].(float64)

		return uint(id)
	}

	return 0
}

func JWTSign() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(config.GetConfig().Secret_JWT),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}
