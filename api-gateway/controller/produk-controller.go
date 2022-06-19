package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/helper"
	"github.com/labstack/echo/v4"
)

type ProdukController interface {
	Insert(ctx echo.Context)
	FindById(ctx echo.Context)
	FindAll(ctx echo.Context)
	Update(ctx echo.Context)
	Delete(ctx echo.Context)
}
type ProdukControllerImpl struct {
}

func NewProdukController() ProdukController {
	return &ProdukControllerImpl{}
}

func (controller *ProdukControllerImpl) FindAll(ctx echo.Context) {
	// authHeader := ctx.Request().Header["Authorization"][0]
	// auth := helper.ReadDataToken(authHeader)
	getall := entity.ReqList{}
	err := ctx.Bind(&getall)
	fmt.Println(err, getall, "NIDDDDD")
	helper.PanicIfError(err)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUK_SERVICE_HOST"))
	// dataRequest := fmt.Sprintf(`{"username":"%s","role":"%s"}`, auth.Username, auth.Role)
	// requestBody := strings.NewReader(dataRequest)
	url := baseURL + fmt.Sprintf("/produk?page=%d&perpage=%d&filter=%s&order=%s", getall.Page, getall.Perpage, getall.Filter, getall.Order)
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

func (controller *ProdukControllerImpl) FindById(ctx echo.Context) {
	// authHeader := ctx.Request().Header["Authorization"][0]
	// auth := helper.ReadDataToken(authHeader)
	getall := entity.ReqListById{}
	err := ctx.Bind(&getall)
	fmt.Println(err, getall, "NIDDDDD")
	helper.PanicIfError(err)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUK_SERVICE_HOST"))
	// dataRequest := fmt.Sprintf(`{"username":"%s","role":"%s"}`, auth.Username, auth.Role)
	// requestBody := strings.NewReader(dataRequest)
	url := baseURL + fmt.Sprintf("/produk/detail?id=%d", getall.Id)
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

func (service *ProdukControllerImpl) Insert(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	gantiPassword := entity.CreateProdukEntity{}
	helper.ReadFromRequestBody(ctx, &gantiPassword)
	fmt.Println(gantiPassword)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUK_SERVICE_HOST"))
	dataRequest := fmt.Sprintf(`{"nama":"%s","kategori":"%s","stok":%d,"username":"%s","role":"%s"}`, gantiPassword.Nama, gantiPassword.Kategori, gantiPassword.Stok, auth.Username, auth.Role)
	requestBody := strings.NewReader(dataRequest)
	fmt.Println(requestBody, "nidzam")
	res, err := http.Post(baseURL+"/produk/insert", "application/json", requestBody)
	fmt.Println(res)
	if err != nil {
		helper.WriteToResponseBody(ctx, err, res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)
	var p entity.WebResponse
	fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	if err != nil {

		helper.WriteToResponseBody(ctx, err, res.StatusCode)
	}
	helper.WriteToResponseBody(ctx, p, p.Code)

}

func (service *ProdukControllerImpl) Update(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	gantiPassword := entity.CreateProdukEntity{}
	helper.ReadFromRequestBody(ctx, &gantiPassword)
	fmt.Println(gantiPassword)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUK_SERVICE_HOST"))
	dataRequest := fmt.Sprintf(`{"id":%d,"nama":"%s","kategori":"%s","stok":%d,"username":"%s","role":"%s"}`, gantiPassword.RowId, gantiPassword.Nama, gantiPassword.Kategori, gantiPassword.Stok, auth.Username, auth.Role)
	requestBody := strings.NewReader(dataRequest)
	fmt.Println(requestBody, "nidzam")
	res, err := http.Post(baseURL+"/produk/update", "application/json", requestBody)
	fmt.Println(res)
	if err != nil {
		helper.WriteToResponseBody(ctx, err, res.StatusCode)
	}

	dec := json.NewDecoder(res.Body)
	var p entity.WebResponse
	fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	if err != nil {

		helper.WriteToResponseBody(ctx, err, res.StatusCode)
	}
	helper.WriteToResponseBody(ctx, p, p.Code)

}

func (controller *ProdukControllerImpl) Delete(ctx echo.Context) {
	authHeader := ctx.Request().Header["Authorization"][0]
	auth := helper.ReadDataToken(authHeader)
	getall := entity.ReqListById{}
	err := ctx.Bind(&getall)
	fmt.Println(err, getall, "NIDDDDD")
	helper.PanicIfError(err)
	baseURL := fmt.Sprintf("http://%s", os.Getenv("PRODUK_SERVICE_HOST"))
	// dataRequest := fmt.Sprintf(`{"username":"%s","role":"%s"}`, auth.Username, auth.Role)
	// requestBody := strings.NewReader(dataRequest)
	url := baseURL + fmt.Sprintf("/produk/delete?id=%d&role=%s", getall.Id, auth.Role)
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
