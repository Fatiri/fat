package main

import (
	"flag"
	"fmt"

	"github.com/Fatiri/fat/app/api"
	"github.com/Fatiri/fat/app/cli"
	"github.com/Fatiri/fat/app/gui"
	"github.com/Fatiri/fat/common/times"
	"github.com/Fatiri/fat/config"
	"github.com/Fatiri/fat/models"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	environment := flag.String("environment", "", "The environment file name")
	service := flag.String("service", "", "The environment file name")
	flag.Parse()

	env := config.NewEnvironment(*environment, "config")
	envData, _ := env.InitEnvironment()
	storage := config.NewStorage(envData)

	config, err := config.NewConfig(env, storage).InitConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	config.Time = times.ProvideNewTimesCustom()

	if *service == "GUI" {
		GUI(config)
	} else if *service == "CLI" {
		CLI(config)
	} else if *service == "API" {
		API(config)
	} else {
		fmt.Println("error : service must be GUI, CLI or API")
	}
}

func GUI(config *models.Config) {
	gui.NewGUI(config).Run()
}

func CLI(config *models.Config) {
	cli.NewCLI(config).Run()
}

func API(config *models.Config) {
	api.NewServer(config).Start()
}
