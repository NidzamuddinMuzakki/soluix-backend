package app

import (
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Nidzam(nid interface{}) error {
	fmt.Println(nid)

	return nil
}
func InitRouter(UserController controller.UserController) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())
	// r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)

	USER := r.Group("user")
	{
		USER.GET("", func(c echo.Context) error {
			UserController.FindAll(c)
			return nil
		})
		USER.GET("/detail", func(c echo.Context) error {
			UserController.FindByUsername(c)
			return nil
		})
		USER.PUT("/detail", func(c echo.Context) error {
			UserController.Insert(c)
			return nil
		})
		USER.POST("/detail", func(c echo.Context) error {
			UserController.Update(c)
			return nil
		})
		USER.POST("/register", func(c echo.Context) error {
			UserController.Register(c)
			return nil
		})

		USER.DELETE("/detail", func(c echo.Context) error {
			UserController.Delete(c)
			return nil
		})
		USER.POST("/login", func(c echo.Context) error {
			UserController.Login(c)
			return nil
		})
	}
	return r
}
