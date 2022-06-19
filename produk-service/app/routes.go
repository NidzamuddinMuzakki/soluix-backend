package app

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(ProdukController controller.UserController) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())
	// r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)
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
		PRODUK.GET("/delete", func(c echo.Context) error {
			ProdukController.Delete(c)
			return nil
		})

	}
	return r
}
