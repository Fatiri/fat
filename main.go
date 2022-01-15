package main

import (
	"flag"
	"fmt"

	"github.com/FAT/api"
	"github.com/FAT/common/times"
	"github.com/FAT/common/wrapper"
	"github.com/FAT/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	environment := flag.String("environment", "", "The environment file name")
	flag.Parse()

	env := config.NewEnvironment(*environment, "./config")
	envData, _ := env.InitEnvironment()
	storage := config.NewStorage(envData)
	
	config, err := config.NewConfig(env, storage).InitConfig()
	if err != nil {
		fmt.Println(wrapper.Error(err, config.Env.EnvApp))
	}
	
	config.Time = times.ProvideNewTimesCustom()

	server := api.NewServer(config)
	server.Start()
}
