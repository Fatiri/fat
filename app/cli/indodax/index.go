package indodax

import (
	"context"
	"fmt"
	"time"

	"github.com/fat/common/indikators"
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
}

func NewIndodaxCLI(config *models.Config, exchange exchange.Indodax, telegram telegram.Telegram) IndodaxCLI {
	return &IndodaxCLICtx{
		config:   config,
		exchange: exchange,
		telegram: telegram,
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
				fmt.Sprintf("Information  \n Saldo : %2.f \n BTC   : %f \n Pair    : %s", saldo, btc, payload.Symbol),
				fmt.Sprintf("RSI   \n Up Price       : %t \n Down Price  : %t \n Close Price  : %2.f \n rsi                 : %2.f", upPrice, downPrice, mh.Close[len(mh.Close)-1], rsiResult),
			}

			go icc.telegram.Bot(messages)
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
