package models

import (
	"github.com/fat/common/times"
	"github.com/fat/repository"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Env         *Environment
	Storage     *repository.Queries
	GinRouter   *gin.Engine
	ServiceType string
	Time        times.Time
}

type Environment struct {
	EnvApp                  string `mapstructure:"ENV_APP"`
	AddressApp              string `mapstructure:"ADDRESS_APP"`
	TitleApp                string `mapstructure:"TITLE_APP"`
	DescriptionApp          string `mapstructure:"DESCRIPTION_APP"`
	VersionApp              string `mapstructure:"VERSION_APP"`
	SchemasApp              string `mapstructure:"SCHEMAS_APP"`
	SymmetricKey            string `mapstructure:"SYMMETRIC_KEY"`
	DatabaseHost            string `mapstructure:"DATABASE_HOST"`
	DatabasePort            string `mapstructure:"DATABASE_PORT"`
	DatabaseUser            string `mapstructure:"DATABASE_USER"`
	DatabasePass            string `mapstructure:"DATABASE_PASS"`
	DatabaseName            string `mapstructure:"DATABASE_NAME"`
	IndodaxPrivateURL       string `mapstructure:"INDODAX_PRIVATE_URL"`
	IndodaxMarketHistoryURL string `mapstructure:"INDODAX_MARKET_HISTORY_URL"`
	IndodaxMarketPendingURL string `mapstructure:"INDODAX_MARKET_PENDING_URL"`
	IndodaxPublicKey        string `mapstructure:"INDODAX_PUBLIC_KEY"`
	IndodaxPrivateKey       string `mapstructure:"INDODAX_PRIVATE_KEY"`
}
