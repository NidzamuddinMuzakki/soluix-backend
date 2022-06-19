package app

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(UserController controller.UserController) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())
	r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)
	USER := r.Group("user")
	{
		USER.POST("/login", func(c echo.Context) error {
			UserController.Login(c)
			return nil
		})
		USER.GET("", func(c echo.Context) error {
			UserController.FindAll(c)
			return nil
		})

	}
	return r
}
