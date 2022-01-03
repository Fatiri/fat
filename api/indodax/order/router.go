package orderapi

import (
	"github.com/fat/middleware"
	"github.com/fat/models"
)

type Router struct {
	orderHandler *OrderHandler
	config       *models.Config
}

func NewRouter(orderHandler *OrderHandler, config *models.Config) *Router {
	return &Router{
		orderHandler: orderHandler,
		config:       config,
	}
}

func (r *Router) RegisterRouter() {
	auth, _ := middleware.NewMaker(r.config.Env.SymmetricKey)
	group := r.config.GinRouter.Group("indodax")
	
	group.POST("/order",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.orderHandler.CreateOrder)
	group.GET("/order/:id", r.orderHandler.GetOrder)
	group.GET("/order/data", r.orderHandler.ListOrder)
}
