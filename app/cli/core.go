package cli

import (
	"fmt"

	"github.com/fat/models"
)

type CLICore interface {
	Run()
}

type CLICoreCtx struct {
	config *models.Config
}

func NewCLI(config *models.Config) CLICore {
	return &CLICoreCtx{config: config}
}

func (cli *CLICoreCtx) init() {
	cli.config.Storage = nil
	fmt.Println("masuk")
}

func (cli *CLICoreCtx) Run() {
	cli.init()
}
