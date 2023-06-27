package indodaxmarketv1

import (
	"github.com/Fatiri/fat/middleware"
	"github.com/Fatiri/fat/models"
)

type Router struct {
	marketHandler *MarketHandler
	config        *models.Config
}

func NewRouter(marketHandler *MarketHandler, config *models.Config) *Router {
	return &Router{
		marketHandler: marketHandler,
		config:        config,
	}
}

func (r *Router) RegisterRouter() {
	auth, _ := middleware.NewAuthentication(r.config)
	indodax := r.config.GinRouter.Group(r.config.ServiceType)

	indodax.POST("/market_history",
		auth.AuthMiddleware([]middleware.Role{middleware.RoleAdmin, middleware.RoleUser}),
		r.marketHandler.MarketHistory)
}
