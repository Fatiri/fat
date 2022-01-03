package config

import (
	"github.com/fat/models"
	"github.com/spf13/viper"
)

func InitEnvironment(envName string) (env models.Environment) {
	viper.AddConfigPath(".")
	viper.SetConfigName(envName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}
	return
}
