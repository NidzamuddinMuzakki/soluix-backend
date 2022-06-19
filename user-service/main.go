package main

import (
	"database/sql"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/app"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/controller"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/repository"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/service"
)

var (
	db             *sql.DB                   = app.Init()
	UserRepository repository.UserRepository = repository.NewUserRepository()
	UserService    service.UserService       = service.NewUserService(UserRepository, db)
	JwtService     service.JWTService        = service.NewJWTService()
	UserController controller.UserController = controller.NewUserController(UserService, JwtService)
)

func main() {
	defer db.Close()

	r := app.InitRouter(UserController)
	r.Start(":9001")
}
