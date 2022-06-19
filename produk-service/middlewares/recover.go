package middlewares

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/exception"
	"github.com/labstack/echo/v4"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				exception.ErrorHandler(c, r)
			}
		}()
		return next(c)

	}
}
