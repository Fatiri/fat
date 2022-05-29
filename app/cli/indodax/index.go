package indodax

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fat/common/constant"
	"github.com/fat/common/indikators"
	"github.com/fat/models"
	"github.com/fat/usecase/exchange"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IndodaxCLI interface {
	Run() error
}

type IndodaxCLICtx struct {
	config   *models.Config
	exchange exchange.Indodax
}

func NewIndodaxCLI(config *models.Config, exchange exchange.Indodax) IndodaxCLI {
	return &IndodaxCLICtx{
		config:   config,
		exchange: exchange,
	}
}

func (icc *IndodaxCLICtx) Run() error {
	payload := models.MarketHistoryPayload{
		Symbol:    "btcidr",
		TimeFrame: "1",
		From:      icc.config.Time.Now(nil).Add(time.Minute * -time.Duration(1440)).Unix(),
		To:        icc.config.Time.Now(nil).Unix(),
	}

	saldo := 15630000.00
	btc := 0.00
	var closePrice, rsiValue float64
	for {
		_, _, second := time.Now().Clock()
		if second == 59 {
			marketHistoryIndodax, errorMarketHistory := icc.exchange.MarketHistory(context.Background(), payload)
			if errorMarketHistory != nil {
				return errorMarketHistory
			}

			mh := new(models.MarketHistory)
			for _, market := range marketHistoryIndodax {
				mh.Close = append(mh.Close, float64(market.Close))
			}
			closePrice = mh.Close[len(mh.Close)-1]

			upPrice, downPrice, rsiResult := icc.RSI(mh.Close)
			if upPrice {
				if saldo == 0 {
					btc = 0
					saldo = mh.Close[len(mh.Close)-1] * float64(btc)
				}
			}

			rsiValue = rsiResult

			if downPrice {
				if btc == 0 {
					btc = float64(saldo) / mh.Close[len(mh.Close)-1]
					saldo = 0
				}
			} 

			fmt.Println("up   : ", upPrice)
			fmt.Println("down : ", downPrice)
			fmt.Println("saldo : ", saldo)
			fmt.Println("btc : ", btc)
			go icc.Telegram(saldo, btc, closePrice, rsiValue, payload.Symbol, upPrice, downPrice)
			time.Sleep(time.Second * 10)
		}
	}
}

func (icc *IndodaxCLICtx) RSI(closePrices []float64) (upPrice bool, downPrice bool, rsiResult float64) {
	rsi := indikators.RsiArr(closePrices, 14)
	if rsi[len(rsi)-1] >= 69 {
		upPrice = true
	}

	if rsi[len(rsi)-1] < 25 {
		downPrice = true
	}

	rsiResult = rsi[len(rsi)-1]

	fmt.Println("-------------------")
	fmt.Println(closePrices[len(closePrices)-1])
	fmt.Println(rsi[len(rsi)-1])

	return upPrice, downPrice, rsiResult
}

func (icc *IndodaxCLICtx) Telegram(saldo, btc, closePrice, rsiValue float64, symbol string, upPrice, DownPrice bool) {
	bot, err := tgbotapi.NewBotAPI("5393311612:AAEwuHHoIAUwxRrhm2YVqad8CWO2V7RB3EQ")
	if err != nil {
		fmt.Println(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	if icc.config.Env.EnvApp == constant.Staging {
		bot.Debug = true
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		msg.Text = fmt.Sprintf("Information  \n Saldo : %2.f \n BTC   : %f \n Pair    : %s", saldo, btc, symbol)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		msg.Text = fmt.Sprintf("RSI   \n Up Price       : %t \n Down Price  : %t \n Close Price  : %2.f \n rsi                 : %2.f", upPrice, DownPrice, closePrice, rsiValue)
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		continue
	}
}
