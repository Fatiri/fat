package api

import (
	"net/http"
	"strings"

	indodaxmarketv1 "github.com/Fatiri/fat/app/api/v1/indodax/market"
	indodaxorderv1 "github.com/Fatiri/fat/app/api/v1/indodax/order"
	"github.com/Fatiri/fat/common/wrapper"
	"github.com/Fatiri/fat/docs"
	"github.com/Fatiri/fat/models"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

type Server interface {
	Start()
}

type ServerCtx struct {
	config *models.Config
}

func NewServer(config *models.Config) Server {
	return &ServerCtx{
		config: config,
	}
}

func (s *ServerCtx) Handler() {
	s.RouterIndodax()
	s.RouterSwagger()

	s.config.GinRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, wrapper.RouteNotFound())
	})
	err := s.config.GinRouter.Run(s.config.Env.AddressApp)
	if err != nil {
		panic(err)
	}
}

func (s *ServerCtx) RouterIndodax() {
	s.config.ServiceType = "indodax"
	indodaxorderv1.InjectOrderRouter(s.config)
	indodaxmarketv1.InjectMarketRouter(s.config)
}

func (s *ServerCtx) RouterSwagger() {
	docs.SwaggerInfo.Title = s.config.Env.TitleApp
	docs.SwaggerInfo.Description = s.config.Env.DescriptionApp
	docs.SwaggerInfo.Version = s.config.Env.VersionApp
	docs.SwaggerInfo.Host = s.config.Env.AddressApp
	docs.SwaggerInfo.Schemes = strings.Split(s.config.Env.SchemasApp, ",")
	s.config.GinRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (s *ServerCtx) Start() {
	s.Handler()
}
