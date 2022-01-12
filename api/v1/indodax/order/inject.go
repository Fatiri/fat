package indodaxorderv1

import "github.com/FAT/models"

func InjectOrderRouter(config *models.Config) {
	handlers := NewOrderHandler(config)
	routers := NewRouter(handlers, config)
	routers.RegisterRouter()
}
