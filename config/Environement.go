package config

import (
	"github.com/FAT/models"
	"github.com/spf13/viper"
)

type Environment interface {
	InitEnvironment() (env *models.Environment, err error)
}

type EnvironmentCtx struct {
	envName string
	dirFile string
}

func NewEnvironment(envName, dirFile string) Environment {
	return &EnvironmentCtx{
		envName: envName,
		dirFile: dirFile,
	}
}

func (e *EnvironmentCtx) InitEnvironment() (env *models.Environment, err error) {
	viper.AddConfigPath(e.dirFile)
	viper.SetConfigName(e.envName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	return
}
