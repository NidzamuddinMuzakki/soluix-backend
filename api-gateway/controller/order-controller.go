package controller

import (
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/entity"
	"github.com/NidzamuddinMuzakki/nidzam-ecomerce/api-gateway/helper"
	"github.com/labstack/echo/v4"
)

type OrderController interface {
	GetData(ctx echo.Context)
}

type OrderControllerImpl struct {
}

func NewOrderController() OrderController {
	return &OrderControllerImpl{}
}

func (contoller *OrderControllerImpl) GetData(ctx echo.Context) {
	webResponse := entity.WebResponseListAndDetail{
		Code: 200,
		Data: "Hay",
		Info: "",
	}
	helper.WriteToResponseBody(ctx, webResponse, webResponse.Code)
}
