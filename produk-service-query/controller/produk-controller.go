package controller

import (
	// 	"fmt"

	// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/entity"
	"encoding/json"

	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/helper"

	// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/rabbitmq"
	// 	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/produk-service-query/service"

	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProdukSearchController interface {
	// FindAll(ctx echo.Context)
	FindSearch(ctx echo.Context)
}

type ProdukSearchControllerImpl struct {
}

func NewProdukSearchController() ProdukSearchController {
	return &ProdukSearchControllerImpl{}
}

type Search struct {
	Words string `query:"search"`
}
type Res struct {
	Hits interface{} `json:"hits"`
	Aggs interface{} `json:"aggregations"`
}

func (controller *ProdukSearchControllerImpl) FindSearch(ctx echo.Context) {
	var search Search
	err := ctx.Bind(&search)
	if err != nil {
		panic(err)
	}
	fmt.Println(search.Words)
	client := &http.Client{}
	url := fmt.Sprintf("http://%s/produk-v1/_search", os.Getenv("ELASTICSEARCH_URL"))
	fmt.Println(url)
	stringRequest := fmt.Sprintf(`{
		"query": {
			"bool": {
				"should": [
								{
									"wildcard": {
										"nama": {
											"value": "*%s*"
										}
									}
								},
								{
									"wildcard": {
										"kategori": {
											"value": "*%s*"
										}
									}
								}
				]
			}
		},
		"aggs": {
			"kategoris": {
				"terms": {
					"field": "kategori.keyword"
				}
			}
		}
	}`, search.Words, search.Words)
	requestBody := strings.NewReader(stringRequest)

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	res, errres := client.Do(req)
	if errres != nil {
		panic(err)
	}
	dec := json.NewDecoder(res.Body)
	var p Res
	// fmt.Println(dec, res, res.Body, p, "nidzam")
	err = dec.Decode(&p)
	fmt.Println(p)
	helper.WriteToResponseBody(ctx, p, res.StatusCode)

}

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
