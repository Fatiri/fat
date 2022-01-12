package config

import (
	"github.com/FAT/models"
	"github.com/gin-gonic/gin"
)

func NewConfig(envName string) *models.Config {
	env := InitEnvironment(envName)
	return &models.Config{
		Env:       env,
		Storage:   Postgres(env),
		GinRouter: gin.Default(),
	}
}
