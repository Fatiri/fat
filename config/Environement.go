package config

import (
	"errors"

	"github.com/Fatiri/fat/models"
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
	if e.envName == "dev" || e.envName == "staging" || e.envName == "production" {
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
	} else {
		return nil, errors.New("error : environment must be dev, staging or production")
	}

	return
}
