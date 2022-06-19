package main

import (
	"database/sql"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/app"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/repository"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/service"
)

var (
	db             *sql.DB                   = app.Init()
	UserRepository repository.UserRepository = repository.NewUserRepository()
	UserService    service.UserService       = service.NewUserService(UserRepository, db)
	UserController controller.UserController = controller.NewUserController(UserService)
)

func main() {
	defer db.Close()

	r := app.InitRouter(UserController)
	r.Start(":9002")
}
