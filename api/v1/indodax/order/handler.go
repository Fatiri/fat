package indodaxorderv1

import (
	"fmt"
	"net/http"

	"github.com/FAT/common/wrapper"
	"github.com/FAT/models"
	"github.com/FAT/repository"
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

// CreateOrder godoc
// @Summary      Create order data
// @Description  Create new order data to database
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string            true  "Authentication header"
// @Param        message        body      repository.Order  true  "Order payload"
// @Success      200            {object}  wrapper.MessageSuccess
// @Failure      400            {object}  wrapper.MessageError
// @Failure      422            {object}  wrapper.MessageError
// @Failure      404            {object}  wrapper.MessageError
// @Failure      401            {object}  wrapper.MessageError
// @Router       /indodax/order [post]
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

// GetOrder godoc
// @Summary      Show an order data
// @Description  get order by ID
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        id             path      int     true  "Order ID"
// @Success      200            {object}  repository.Order
// @Failure      400            {object}  wrapper.MessageError
// @Failure      404            {object}  wrapper.MessageError
// @Failure      401            {object}  wrapper.MessageError
// @Router       /indodax/order/:id [get]
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

// ListOrder godoc
// @Summary      List order
// @Description  get orders
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param        page    query  int32  false "page offset"
// @Param        limit   query   int32  false "data size"
// @Success      200  {array}  repository.Order
// @Failure      400  {object}  wrapper.MessageError
// @Failure      404  {object}  wrapper.MessageError
// @Failure      401  {object}  wrapper.MessageError
// @Failure      500  {object}  wrapper.MessageError
// @Router       /indodax/order/data [get]
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

// DeleteOrder godoc
// @Summary      Delete order
// @Description  Delete order Data by ID
// @Param        Authorization  header  string  true  "Authentication header"
// @Param        id  path  int  true  "Order ID"
// @Success      200  {object}  wrapper.MessageSuccess
// @Failure      400  {object}  wrapper.MessageError
// @Failure      404  {object}  wrapper.MessageError
// @Failure      500  {object}  wrapper.MessageError
// @Failure      401  {object}  wrapper.MessageError
// @Router       /indodax/order/:id [delete]
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

	ctx.JSON(http.StatusOK, wrapper.SuccessHandler("success"))
}
