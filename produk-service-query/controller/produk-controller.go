package controller

// import (
// 	"fmt"

// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/entity"
// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/helper"

// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/rabbitmq"
// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/service"

// 	"github.com/labstack/echo/v4"
// )

// type UserController interface {
// 	FindAll(ctx echo.Context)
// 	FindById(ctx echo.Context)
// }
// type LoginDTO struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Role     string `json:"role"`
// }
// type LoginResponse struct {
// 	Token        string `json:"token"`
// 	RefreshToken string `json:"refresh_token"`
// }

// type UserControllerImpl struct {
// 	UserService service.UserService
// 	Rc          rabbitmq.RabbitClient
// }

// func NewUserController(userService service.UserService, rc rabbitmq.RabbitClient) UserController {
// 	return &UserControllerImpl{
// 		UserService: userService,
// 		Rc:          rc,
// 	}
// }

// type funcName struct {
// 	Nama string `json:"nama"`
// }

// func (controller *UserControllerImpl) FindAll(ctx echo.Context) {

// 	getall := entity.ReqList{}
// 	err := ctx.Bind(&getall)

// 	helper.PanicIfError(err)

// 	resultData := controller.UserService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order)
// 	fmt.Println(resultData)

// 	webResponse := entity.WebResponseListAndDetail{
// 		Code: 200,
// 		Data: resultData,
// 		Info: "",
// 	}

// 	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
// }

// func (controller *UserControllerImpl) FindById(ctx echo.Context) {
// 	// authHeader := ctx.Request().Header["Authorization"][0]
// 	// auth := helper.ReadDataToken(authHeader)
// 	// fmt.Println(auth)
// 	getall := entity.ReqListById{}
// 	err := ctx.Bind(&getall)
// 	// fmt.Println(err, getall.Username)
// 	helper.PanicIfError(err)

// 	resultData := controller.UserService.FindById(ctx.Request().Context(), getall.Id)
// 	// fmt.Println(beli)
// 	webResponse := entity.WebResponseListAndDetail{
// 		Code: 200,
// 		Data: resultData,
// 		Info: "",
// 	}

// 	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
// }
