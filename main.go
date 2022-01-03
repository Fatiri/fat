package main

import (
	"flag"

	"github.com/fat/api"
	"github.com/fat/config"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	environment := flag.String("environment", "", "The environment file name")
	flag.Parse()

	config := config.NewConfig(*environment)
	server := api.NewServer(config)
	server.Start()
}
