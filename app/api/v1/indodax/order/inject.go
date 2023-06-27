package indodaxorderv1

import "github.com/Fatiri/fat/models"

func InjectOrderRouter(config *models.Config) {
	handlers := NewOrderHandler(config)
	routers := NewRouter(handlers, config)
	routers.RegisterRouter()
}
