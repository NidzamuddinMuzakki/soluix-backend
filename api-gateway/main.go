package main

import (
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/app"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/service"
)

var (
	userJWTService service.JWTService        = service.NewJWTService()
	UserController controller.UserController = controller.NewUserController(userJWTService)
)

func main() {
	fmt.Println("nidzam")

	r := app.InitRouter(UserController)
	r.Start(":9000")
}
