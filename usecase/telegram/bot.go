package telegram

import (
	"fmt"
	"log"

	"github.com/fat/common/constant"
	"github.com/fat/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram interface {
	Bot(messages []string)
}

type TelegramCtx struct {
	config *models.Config
}

func NewTelegram(config *models.Config) Telegram {
	return &TelegramCtx{
		config: config,
	}
}

func (tc *TelegramCtx) Bot(messages []string) {
	bot, err := tgbotapi.NewBotAPI(tc.config.Env.TelegramAPIToken)
	if err != nil {
		fmt.Println(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	if tc.config.Env.EnvApp == constant.Staging {
		bot.Debug = true
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.

		for _, message := range messages {
			msg.Text = message
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
