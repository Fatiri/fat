package models

import (
	"github.com/FAT/common/times"
	"github.com/FAT/repository"
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
	EnvApp            string `mapstructure:"ENV_APP"`
	AddressApp        string `mapstructure:"ADDRESS_APP"`
	TitleApp          string `mapstructure:"TITLE_APP"`
	DescriptionApp    string `mapstructure:"DESCRIPTION_APP"`
	VersionApp        string `mapstructure:"VERSION_APP"`
	SchemasApp        string `mapstructure:"SCHEMAS_APP"`
	SymmetricKey      string `mapstructure:"SYMMETRIC_KEY"`
	DatabaseHost      string `mapstructure:"DATABASE_HOST"`
	DatabasePort      string `mapstructure:"DATABASE_PORT"`
	DatabaseUser      string `mapstructure:"DATABASE_USER"`
	DatabasePass      string `mapstructure:"DATABASE_PASS"`
	DatabaseName      string `mapstructure:"DATABASE_NAME"`
	IndodaxPrivateURL string `mapstructure:"INDODAX_PRIVATE_URL"`
}
