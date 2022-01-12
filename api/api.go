package api

import (
	"net/http"
	"strings"

	indodaxorderv1 "github.com/FAT/api/v1/indodax_order"
	"github.com/FAT/common/wrapper"
	"github.com/FAT/docs"
	"github.com/FAT/models"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	s.config.GinRouter = gin.Default()

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
