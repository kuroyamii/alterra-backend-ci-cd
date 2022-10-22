package middleware

import (
	"belajar-go-echo/pkg/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Authorization"] != nil {
			auth := strings.Split(c.Request().Header["Authorization"][0], " ")
			token, err := jwt.Parse(auth[1], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
						"error": "ACCESS_DENIED",
					})
				}
				return []byte(config.JWT_KEY), nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
					"error": "ACCESS_DENIED",
				})
			}

			if token.Valid {
				return next(c)
			}
			return nil
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, echo.Map{
				"error": "ACCESS_DENIED",
			})
		}
	}
}
