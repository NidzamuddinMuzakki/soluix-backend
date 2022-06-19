package controller

import (
	"fmt"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/helper"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/user-service/service"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	FindAll(ctx echo.Context)
	FindByUsername(ctx echo.Context)
	Insert(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
	Login(ctx echo.Context)
	Register(ctx echo.Context)
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
	JwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &UserControllerImpl{
		UserService: userService,
		JwtService:  jwtService,
	}
}

func (service *UserControllerImpl) Login(ctx echo.Context) {
	loginDTO := LoginDTO{}
	helper.ReadFromRequestBody(ctx, &loginDTO)

	authResult, Role := service.UserService.VerifyCredential(ctx.Request().Context(), loginDTO.Username, loginDTO.Password)
	if authResult == true {
		fmt.Println(loginDTO.Username, Role, "nidzam")
		// generatedToken, refresh_token := service.JwtService.GenerateToken(loginDTO.Username, Role)
		loginDTO.Role = Role
		webResponse := entity.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   loginDTO,
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	} else {
		webResponse := entity.WebResponse{
			Code:   401,
			Status: "unauthorized",
			Data:   "gagal",
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	}
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	fmt.Println(authHeader)
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

func (controller *UserControllerImpl) FindByUsername(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	fmt.Println(auth)
	getall := entity.ReqListByUsername{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)

	resultData := controller.UserService.FindByUsername(ctx.Request().Context(), getall.Username)
	// fmt.Println(beli)
	webResponse := entity.WebResponseListAndDetail{
		Code: 200,
		Data: resultData,
		Info: "",
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Insert(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)

	CreateRequest := entity.UserEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	resultData := controller.UserService.Insert(ctx.Request().Context(), CreateRequest, auth.Username, auth.Role)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Register(ctx echo.Context) {

	CreateRequest := entity.UserEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	resultData := controller.UserService.Register(ctx.Request().Context(), CreateRequest)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
func (controller *UserControllerImpl) Update(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)

	CreateRequest := entity.UserEntity{}

	helper.ReadFromRequestBody(ctx, &CreateRequest)
	resultData := controller.UserService.Update(ctx.Request().Context(), CreateRequest, auth.Username, auth.Role)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}

func (controller *UserControllerImpl) Delete(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	getall := entity.ReqListByUsername{}
	err := ctx.Bind(&getall)

	helper.PanicIfError(err)
	resultData := controller.UserService.Delete(ctx.Request().Context(), getall.Username, auth.Role)
	webResponse := entity.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resultData,
	}

	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
