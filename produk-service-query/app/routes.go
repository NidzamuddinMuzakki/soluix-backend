package app

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouter(ProdukSearchController controller.ProdukSearchController) *echo.Echo {

	r := echo.New()
	r.Use(middleware.CORS())

	// r.Use(middlewares.Auth)
	r.Use(middlewares.Recover)
	PRODUK := r.Group("produk")
	{
		PRODUK.GET("/search", func(c echo.Context) error {
			ProdukSearchController.FindSearch(c)
			return nil
		})

	}
	return r
}
