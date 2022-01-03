package api

import (
	"net/http"

	orderapi "github.com/fat/api/indodax/order"
	"github.com/fat/common/wrapper"
	"github.com/fat/models"
	"github.com/gin-gonic/gin"
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
	s.config.GinRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, wrapper.RouteNotFound())
	})
	err := s.config.GinRouter.Run(s.config.Env.AddressApp)
	if err != nil {
		panic(err)
	}
}

func (s *ServerCtx) RouterIndodax() {
	orderapi.InjectOrderRouter(s.config)
}

func (s *ServerCtx) Start() {
	s.Handler()
}
