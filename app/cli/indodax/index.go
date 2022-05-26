package indodax

import (
	"context"
	"fmt"
	"time"

	"github.com/fat/common/indikators"
	"github.com/fat/models"
	"github.com/fat/usecase/exchange"
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
	saldo := 200000.00
	btc := 0.00
	for {
		_, _, second := time.Now().Clock()
		if second == 59 {
			payload := models.MarketHistoryPayload{
				Symbol:    "btcidr",
				TimeFrame: "1",
				From:      icc.config.Time.Now(nil).Add(time.Minute * -time.Duration(1440)).Unix(),
				To:        icc.config.Time.Now(nil).Unix(),
			}

			marketHistoryIndodax, errorMarketHistory := icc.exchange.MarketHistory(context.Background(), payload)
			if errorMarketHistory != nil {
				return errorMarketHistory
			}

			mh := new(models.MarketHistory)
			for _, market := range marketHistoryIndodax {
				mh.Close = append(mh.Close, float64(market.Close))
			}

			upPrice, downPrice := icc.RSI(mh.Close)
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
			fmt.Println("up   : ", upPrice)
			fmt.Println("down : ", downPrice)
			fmt.Println("saldo : ", saldo)
			fmt.Println("btc : ", btc)
			time.Sleep(time.Second * 10)
		}
	}
}

func (icc *IndodaxCLICtx) RSI(closePrices []float64) (upPrice bool, downPrice bool) {
	rsi := indikators.RsiArr(closePrices, 14)
	if rsi[len(rsi)-1] >= 69 {
		upPrice = true
	}

	if rsi[len(rsi)-1] <= 30 {
		downPrice = true
	}

	fmt.Println("-------------------")
	fmt.Println(rsi[len(rsi)-1])

	return upPrice, downPrice
}
