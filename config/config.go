package config

import (
	"github.com/fat/models"
	"github.com/gin-gonic/gin"
)

type Config interface {
	InitConfig() (conf *models.Config, err error)
}

type ConfigCtx struct {
	env     Environment
	storage Storage
}

func NewConfig(env Environment, storage Storage) Config {
	return &ConfigCtx{
		env:     env,
		storage: storage,
	}
}

func (c *ConfigCtx) InitConfig() (conf *models.Config, err error) {
	env, err := c.env.InitEnvironment()
	if err != nil {
		return nil, err
	}
	storage, err := c.storage.Postgres()
	if err != nil {
		return nil, err
	}

	conf = &models.Config{
		Env:       env,
		Storage:   storage,
		ServiceType: env.EnvApp,
		GinRouter: gin.Default(),
	}

	return
}
