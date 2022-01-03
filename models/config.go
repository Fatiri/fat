package models

import (
	"github.com/fat/repository"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Env       Environment
	Storage   *repository.Queries
	GinRouter *gin.Engine
	Version   string
}

type Environment struct {
	EnvApp   string `mapstructure:"ENV_APP"`
	AddressApp   string `mapstructure:"ADDRESS_APP"`
	DatabaseHost string `mapstructure:"DATABASE_HOST"`
	DatabasePort string `mapstructure:"DATABASE_PORT"`
	DatabaseUser string `mapstructure:"DATABASE_USER"`
	DatabasePass string `mapstructure:"DATABASE_PASS"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
	SymmetricKey string `mapstructure:"SYMMETRIC_KEY"`
}
