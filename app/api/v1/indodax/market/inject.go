package indodaxmarketv1

import "github.com/Fatiri/fat/models"

func InjectMarketRouter(config *models.Config) {
	handlers := NewMarketHandler(config)
	routers := NewRouter(handlers, config)
	routers.RegisterRouter()
}
