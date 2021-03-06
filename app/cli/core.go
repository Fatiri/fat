package cli

import (
	"time"

	"github.com/fat/app/cli/indodax"
	"github.com/fat/common/mailer"
	"github.com/fat/models"
	"github.com/fat/usecase/exchange"
	"github.com/fat/usecase/telegram"
)

type CLICore interface {
	Run()
}

type CLICoreCtx struct {
	config   *models.Config
	exchange exchange.Indodax
	telegram telegram.Telegram
	gmail    mailer.Gmail
}

func NewCLI(config *models.Config) CLICore {
	return &CLICoreCtx{
		config:   config,
		exchange: exchange.NewIndodax(config),
		telegram: telegram.NewTelegram(config),
		gmail: mailer.NewGmail(mailer.GmailConfig{
			SenderAddress: config.Env.GmailSenderAddress,
			Host:          config.Env.GmailHost,
			Port:          config.Env.GmailPort,
			Username:      config.Env.GmailUsername,
			Password:      config.Env.GmailPassword,
		}),
	}
}

func (cli *CLICoreCtx) init() {
	cli.config.Storage = nil

	indodax.NewIndodaxCLI(cli.config, cli.exchange, cli.telegram, cli.gmail).Run()
	time.Sleep(time.Second * 10)

}

func (cli *CLICoreCtx) Run() {
	cli.init()
}
