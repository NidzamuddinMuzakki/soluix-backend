package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/helper"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/service"

	"github.com/labstack/echo/v4"
)

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseLogin struct {
	Authresult bool
	Role       string
}

type ResponseLoginBody struct {
	Code   int      `json:"code"`
	Status string   `json:"status"`
	Data   LoginDTO `json:"data"`
}

type ResponseLoginBodyCode struct {
	Code         int    `json:"code"`
	Status       string `json:"status"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type UserController interface {
	Login(ctx echo.Context)
	FindAll(ctx echo.Context)
}
type UserControllerImpl struct {
	JwtService service.JWTService
}

func NewUserController(jwtService service.JWTService) UserController {
	return &UserControllerImpl{
		JwtService: jwtService,
	}
}

func (service *UserControllerImpl) Login(ctx echo.Context) {
	loginDTO := LoginDTO{}
	helper.ReadFromRequestBody(ctx, &loginDTO)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("USER_SERVICE_HOST"))
	dataRequest := fmt.Sprintf(`{"username" : "%s","password":"%s"}`, loginDTO.Username, loginDTO.Password)
	requestBody := strings.NewReader(dataRequest)
	res, err := http.Post(baseURL+"/user/login", "application/json", requestBody)
	fmt.Println(res)
	if err != nil {
		webResponse := entity.WebResponse{
			Code:   401,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
	}

	dec := json.NewDecoder(res.Body)
	var p ResponseLoginBody
	fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	if err != nil {
		webResponse := entity.WebResponse{
			Code:   401,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
	}
	if p.Data.Username != "" {
		fmt.Println(loginDTO.Username, p.Data.Role)
		generatedToken, refresh_token := service.JwtService.GenerateToken(loginDTO.Username, p.Data.Role)
		DataLogin := LoginResponse{}
		DataLogin.Token = generatedToken
		DataLogin.RefreshToken = refresh_token
		webResponse := ResponseLoginBodyCode{
			Code:         200,
			Status:       "OK",
			Token:        DataLogin.Token,
			RefreshToken: DataLogin.RefreshToken,
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	} else {
		webResponse := entity.WebResponse{
			Code:   401,
			Status: "gagal",
			Data:   "gagal",
		}

		helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)

	}
}

func (controller *UserControllerImpl) FindAll(ctx echo.Context) {
	// authHeader := ctx.Request().Header["Authorization"][0]
	// auth := helper.ReadDataToken(authHeader)
	getall := entity.ReqList{}
	err := ctx.Bind(&getall)
	fmt.Println(err, getall, "NIDDDDD")
	helper.PanicIfError(err)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("USER_SERVICE_HOST"))
	// dataRequest := fmt.Sprintf(`{"username":"%s","role":"%s"}`, auth.Username, auth.Role)
	// requestBody := strings.NewReader(dataRequest)
	url := baseURL + fmt.Sprintf("/user?page=%d&perpage=%d&filter=%s&order=%s", getall.Page, getall.Perpage, getall.Filter, getall.Order)
	fmt.Println(url, "cek url")
	res, err := http.Get(url)

	fmt.Println(err, res, "nidzazazaza")
	dec := json.NewDecoder(res.Body)
	var p entity.WebResponseListAndDetail
	// fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	fmt.Println(err, p, "NIDZZZ")

	// fmt.Println(beli)
	// webResponse := entity.WebResponseListAndDetail{
	// 	Code: 200,
	// 	Data: resultData,
	// 	Info: "",
	// }

	helper.WriteToResponseBody(ctx, p, p.Code)
}
