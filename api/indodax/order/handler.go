package orderapi

import (
	"fmt"
	"net/http"

	"github.com/fat/common/wrapper"
	"github.com/fat/models"
	"github.com/fat/repository"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	config *models.Config
}

func NewOrderHandler(config *models.Config) *OrderHandler {
	return &OrderHandler{
		config: config,
	}
}

func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req repository.CreateOrderParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	order, err := oh.config.Storage.CreateOrder(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

type getOrderRequest struct {
		OrderID int64 `uri:"id" binding:"required"`
}

func (oh *OrderHandler) GetOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	order, err := oh.config.Storage.GetOrder(ctx, req.OrderID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, wrapper.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, order)
}

type listOrderRequest struct {
	Page  int32 `form:"page" binding:"required,min=1"`
	Limit int32 `form:"limit" binding:"required,min=1"`
}

func (oh *OrderHandler) ListOrder(ctx *gin.Context) {
	var req listOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	arg := repository.ListOrderParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	orders, err := oh.config.Storage.ListOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, wrapper.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (oh *OrderHandler) DeleteOrder(ctx *gin.Context) {
	var req repository.CreateOrderParams
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	err := oh.config.Storage.DeleteOrder(ctx, req.OrderID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.ErrorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, wrapper.SuccessHandler())
}
