package controller

import (
	"encoding/json"
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/helper"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/rabbitmq"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service/service"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	FindAll(ctx echo.Context)
	FindById(ctx echo.Context)
	Insert(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
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

	resultData := controller.UserService.FindAll(ctx.Request().Context(), getall.Page, getall.Perpage, getall.Filter, getall.Order)
	fmt.Println(resultData)

	webResponse := entity.WebResponseListAndDetail{
		Code: 200,
		Data: resultData,
		Info: "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

func (controller *UserControllerImpl) FindById(ctx echo.Context) {
	// authHeader := ctx.Request().Header["Authorization"][0]
	// auth := helper.ReadDataToken(authHeader)
	// fmt.Println(auth)
	getall := entity.ReqListById{}
	err := ctx.Bind(&getall)
	// fmt.Println(err, getall.Username)
	helper.PanicIfError(err)

	resultData := controller.UserService.FindById(ctx.Request().Context(), getall.Id)
	// fmt.Println(beli)
	webResponse := entity.WebResponseListAndDetail{
		Code: 200,
		Data: resultData,
		Info: "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Insert(ctx echo.Context) {

	CreateRequest := entity.ProdukEntity{}

	err := ctx.Bind(&CreateRequest)
	fmt.Println(err)
	helper.PanicIfError(err)
	fmt.Println(CreateRequest, "nidzam")
	resultData, dataFromInsert := controller.UserService.Insert(ctx.Request().Context(), CreateRequest)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}
	newFsConfigBytes, _ := json.Marshal(dataFromInsert)
	go controller.Rc.Publish("insert-queue", newFsConfigBytes)

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

func (controller *UserControllerImpl) Update(ctx echo.Context) {

	CreateRequest := entity.ProdukEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	fmt.Println(CreateRequest, "hayayaasdasas")
	if CreateRequest.Role != "admin" {
		webResponse := entity.WebResponse{
			Code:   406,
			Status: "false",
			// Data:   resultData,
			Data: "anda bukan admin",
		}
		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
		return
	}
	resultData, dataFromUpdate := controller.UserService.Update(ctx.Request().Context(), CreateRequest)
	newFsConfigBytes, _ := json.Marshal(dataFromUpdate)
	go controller.Rc.Publish("update-queue", newFsConfigBytes)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		// Data:   resultData,
		Data: resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

type GETId struct {
	RowId int    `query:"id"`
	Role  string `query:"role"`
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) {

	getall := GETId{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)
	if getall.Role != "admin" {
		webResponse := entity.WebResponse{
			Code:   406,
			Status: "false",
			// Data:   resultData,
			Data: "anda bukan admin",
		}
		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
		return
	}
	resultData := controller.UserService.Delete(ctx.Request().Context(), getall.RowId)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}
	newFsConfigBytes, _ := json.Marshal(getall)
	go controller.Rc.Publish("delete-queue", newFsConfigBytes)

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
