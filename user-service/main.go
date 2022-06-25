package main

import (
	"database/sql"
	"fmt"

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

type funcName struct {
	Nama string `json:"nama"`
}

func Nidzam(nid interface{}) error {

	fmt.Println(nid)
	return nil
}
func main() {
	defer db.Close()
	// var rc rabbitmq.RabbitClient
	// go rc.Consume("test-queue", Nidzam)
	r := app.InitRouter(UserController)
	r.Start(":9001")
}
