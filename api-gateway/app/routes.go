package app

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(UserController controller.UserController, ProdukController controller.ProdukController, OrderController controller.OrderController) *echo.Echo {

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
		USER.GET("/detail", func(c echo.Context) error {
			UserController.FindByUsername(c)
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

	}
	PRODUK := r.Group("produk")
	{

		PRODUK.GET("", func(c echo.Context) error {
			ProdukController.FindAll(c)
			return nil
		})
		PRODUK.GET("/detail", func(c echo.Context) error {
			ProdukController.FindById(c)
			return nil
		})
		PRODUK.POST("/insert", func(c echo.Context) error {
			ProdukController.Insert(c)
			return nil
		})
		PRODUK.POST("/update", func(c echo.Context) error {
			ProdukController.Update(c)
			return nil
		})
		PRODUK.DELETE("/delete", func(c echo.Context) error {
			ProdukController.Delete(c)
			return nil
		})

	}
	ORDER := r.Group("order")
	{
		ORDER.GET("", func(c echo.Context) error {
			OrderController.GetData(c)
			return nil
		})
	}
	return r
}
