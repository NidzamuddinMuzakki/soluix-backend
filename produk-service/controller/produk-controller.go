package controller

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/rabbitmq"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	InsertAwal(ctx echo.Context)
	FindAll(ctx echo.Context)
	// FindById(ctx echo.Context)
}
type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserControllerImpl struct {
	UserService service.UserService
	Rc          rabbitmq.RabbitClient
}

func NewUserController(userService service.UserService, rc rabbitmq.RabbitClient) UserController {
	return &UserControllerImpl{
		UserService: userService,
		Rc:          rc,
	}
}

type funcName struct {
	Nama string `json:"nama"`
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) {

	getall := entity.ReqList{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)

	resultData, info := controller.UserService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order)
	fmt.Println(resultData)
	status := "berhasil"
	code := 200
	if len(resultData) == 0 {
		status = "not found"
		code = 404
	}
	webResponse := entity.WebResponseListAndDetail{
		Code:   code,
		Status: status,
		Data:   resultData,
		Info:   info,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

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

func RandomString(n int) string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
func RandomInt(n int) string {
	var letters = []rune("1234567890")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
func (controller *UserControllerImpl) InsertAwal(ctx echo.Context) {

	var wg sync.WaitGroup
	cuyNow := helper.TimePlus7(time.Now())
	// resultData := controller.UserService.Insert(ctx.Request().Context(), CreateRequest)
	for i := 1; i <= 200000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			CreateRequest := entity.ProdukEntity{}
			CreateRequest.ProductName = fmt.Sprintf("barang%d", i)
			CreateRequest.ProductId = RandomString(3) + "-" + RandomInt(4) + "-" + RandomString(3)
			CreateRequest.SubCategory = "barang baru"
			CreateRequest.Brand = "clarity"
			CreateRequest.Status = "active"
			CreateRequest.CreatedBy = "admin"
			CreateRequest.CreatedTime = cuyNow
			CreateRequest.Price = rand.Intn(50000-10000) + 10000
			controller.UserService.Insert(ctx.Request().Context(), CreateRequest)
		}()
	}
	wg.Wait()
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

type GETId struct {
	RowId int    `query:"id"`
	Role  string `query:"role"`
}
