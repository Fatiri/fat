package indodax

import (
	"context"
	"fmt"
	"time"

	"github.com/fat/common/indikators"
	"github.com/fat/common/mailer"
	"github.com/fat/models"
	"github.com/fat/usecase/exchange"
	"github.com/fat/usecase/telegram"
)

type IndodaxCLI interface {
	Run() error
}

type IndodaxCLICtx struct {
	config   *models.Config
	exchange exchange.Indodax
	telegram telegram.Telegram
	gmail    mailer.Gmail
}

func NewIndodaxCLI(config *models.Config, exchange exchange.Indodax, telegram telegram.Telegram, gmail mailer.Gmail) IndodaxCLI {
	return &IndodaxCLICtx{
		config:   config,
		exchange: exchange,
		telegram: telegram,
		gmail:    gmail,
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
	for {
		hour, minute, second := time.Now().Clock()
		if second == 59 {
			marketHistoryIndodax, errorMarketHistory := icc.exchange.MarketHistory(context.Background(), payload)
			if errorMarketHistory != nil {
				return errorMarketHistory
			}

			mh := new(models.MarketHistory)
			for _, market := range marketHistoryIndodax {
				mh.Close = append(mh.Close, float64(market.Close))
			}

			upPrice, downPrice, rsiResult := icc.RSI(mh.Close)
			if upPrice {
				if saldo == 0 {
					btc = 0
					saldo = mh.Close[len(mh.Close)-1] * float64(btc)
				}
			}

			if downPrice {
				if btc == 0 {
					btc = float64(saldo) / mh.Close[len(mh.Close)-1]
					saldo = 0
				}
			}

			messages := []string{
				fmt.Sprintf("Information  <br> Saldo : %2.f <br> BTC : %f <br> Pair : %s <br><br>", saldo, btc, payload.Symbol),
				fmt.Sprintf("RSI   <br> Up Price : %t <br> Down Price : %t <br> Close Price : %2.f <br> rsi : %2.f", upPrice, downPrice, mh.Close[len(mh.Close)-1], rsiResult),
			}
			fmt.Println(hour,minute)

			if (hour%2 == 0 && minute == 1) {
				icc.gmail.V1(mailer.GmailPayload{
					ReceiverEmail: icc.config.Env.ReceiverMailReport,
					Subject:       icc.config.Env.TitleApp,
					Message:       messages[0] + messages[1],
				})
			}
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

	return upPrice, downPrice, rsiResult
}
