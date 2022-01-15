package indodaxorderv1

import (
	"fmt"

	"github.com/FAT/middleware"
	"github.com/FAT/models"
	"github.com/FAT/repository"
	"github.com/google/uuid"
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
	auth, _ := middleware.NewAuthentication(r.config)
	indodax := r.config.GinRouter.Group(r.config.ServiceType)
	uuid,_ := uuid.NewRandom()
	acc := repository.Account{
		AccountID: uuid,
		AccountType: "ADMIN",
	}
	token, _ := auth.CreateToken(acc)
	fmt.Println(token)

	indodax.POST("/order",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.orderHandler.CreateOrder)
	indodax.GET("/order/:id",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.orderHandler.GetOrder)
	indodax.GET("/order/data",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.orderHandler.ListOrder)
	indodax.DELETE("/order/:id",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.orderHandler.DeleteOrder)
}
