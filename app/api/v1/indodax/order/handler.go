package indodaxorderv1

import (
	"net/http"

	"github.com/Fatiri/fat/common/wrapper"
	"github.com/Fatiri/fat/models"
	"github.com/Fatiri/fat/repository"
	"github.com/Fatiri/fat/usecase/exchange"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	config  *models.Config
	indodax exchange.Indodax
}

func NewOrderHandler(config *models.Config) *OrderHandler {
	return &OrderHandler{
		config:  config,
		indodax: exchange.NewIndodax(config),
	}
}

// CreateOrder godoc
// @Summary      Create order data
// @Description  Create new order data to database
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string            true  "Authentication header"
// @Param        message        body      models.OrderPayload  true  "Order payload"
// @Success      200            {object}  wrapper.Response
// @Failure      400            {object}  wrapper.Response
// @Failure      422            {object}  wrapper.Response
// @Failure      404            {object}  wrapper.Response
// @Failure      401            {object}  wrapper.Response
// @Router       /indodax/order [post]
func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var payload models.OrderPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	order, err := oh.indodax.Order(ctx, repository.CreateOrderParams{
		OrderType:   payload.OrderType,
		OrderPrice:  payload.OrderPrice,
		OrderCrypto: payload.OrderCrypto,
		CreatedAt:   oh.config.Time.Now(nil),
		UpdatedAt:   oh.config.Time.Now(nil),
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, wrapper.Error(err, oh.config.Env.EnvApp))
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
// @Failure      400            {object}  wrapper.Response
// @Failure      404            {object}  wrapper.Response
// @Failure      401            {object}  wrapper.Response
// @Router       /indodax/order/{id} [get]
func (oh *OrderHandler) GetOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	order, err := oh.config.Storage.GetOrder(ctx, req.OrderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, wrapper.Error(err, oh.config.Env.EnvApp))
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
// @Failure      400  {object}  wrapper.Response
// @Failure      404  {object}  wrapper.Response
// @Failure      401  {object}  wrapper.Response
// @Failure      500  {object}  wrapper.Response
// @Router       /indodax/order/data [get]
func (oh *OrderHandler) ListOrder(ctx *gin.Context) {
	var req listOrderRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	arg := repository.ListOrderParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	orders, err := oh.config.Storage.ListOrder(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// DeleteOrder godoc
// @Summary      Delete order
// @Description  Delete order Data by ID
// @Param        Authorization  header  string  true  "Authentication header"
// @Param        id  path  int  true  "Order ID"
// @Success      200  {object}  wrapper.Response
// @Failure      400  {object}  wrapper.Response
// @Failure      404  {object}  wrapper.Response
// @Failure      500  {object}  wrapper.Response
// @Failure      401  {object}  wrapper.Response
// @Router       /indodax/order/{id} [delete]
func (oh *OrderHandler) DeleteOrder(ctx *gin.Context) {
	var req repository.CreateOrderParams
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	err := oh.config.Storage.DeleteOrder(ctx, req.OrderID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, wrapper.Error(err, oh.config.Env.EnvApp))
		return
	}

	ctx.JSON(http.StatusOK, wrapper.Success("success"))
}
