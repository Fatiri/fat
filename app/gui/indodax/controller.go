package indodaxgui

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	"fyne.io/fyne/v2/dialog"
	"github.com/fat/common/indikators"
	"github.com/fat/common/manipulator"
	"github.com/fat/models"
	"github.com/fat/repository"
	"github.com/yudapc/go-rupiah"
)

func (ig *IndodaxGUICtx) BOT() {
	payload := models.MarketHistoryPayload{
		Symbol:    "btcidr",
		TimeFrame: "1",
		From:      ig.config.Time.Now(nil).Add(time.Minute * -time.Duration(1440)).Unix(),
		To:        ig.config.Time.Now(nil).Unix(),
	}

	tm := time.NewTicker(time.Second * 1)
	var isDown, isdownEqual, isUp, isDownRSI, isSell, isBuy bool
	balance := 400000
	var crypto float64
	var price int64
	countDownSell := 0
	for range tm.C {
		_, _, s := ig.config.Time.Now(nil).Clock()
		mHistory, _ := ig.exchange.MarketHistory(context.Background(), payload)
		if len(mHistory) > 0 {
			mh := new(models.MarketHistory)

			for _, market := range mHistory {
				mh.Close = append(mh.Close, float64(market.Close))
			}
			a, b, c := indikators.BBandsArr(0, mh.Close, 20, 2, 2)
			fmt.Println(a[len(a)-1], b[len(b)-1], c[len(c)-1])
			rsi := indikators.RsiArr(mh.Close, 14)
			_, _, g := indikators.MacdArr(mh.Close, 12, 26, 9)

			if s == 1 {
				if rsi[len(rsi)-1] < 30 {
					isDownRSI = true
				} else {
					isDownRSI = false
				}

				if mh.Close[len(mh.Close)-1] > b[len(b)-1] {
					countDownSell++
					isUp = true
				} else {
					countDownSell = 0
					isUp = false
				}

				if mh.Close[len(mh.Close)-1] > c[len(c)-1] && mh.Close[len(mh.Close)-2] < c[len(c)-2] &&
					mh.Close[len(mh.Close)-3] < c[len(c)-3] && mh.Close[len(mh.Close)-4] < c[len(c)-4] {
					isDown = true
				} else {
					isDown = false
				}

				if mh.Close[len(mh.Close)-1] == c[len(c)-1] {
					isdownEqual = true
				} else {
					isdownEqual = false
				}

			}

			fmt.Println(isDownRSI, crypto, balance)

			if s >= 58 {
				if isUp && countDownSell >= 1 {
					isSell = true
					isBuy = false
					price = mHistory[len(mHistory)-1].Close + 10000
				}

				if isDown && g[len(g)-1] < 0 && g[len(g)-1] < g[len(g)-2] && rsi[len(rsi)-1] < 28 {
					isBuy = true
					price = mHistory[len(mHistory)-1].Close - 10000
					isSell = false
				}
			}

			if isBuy && !isSell {
				if crypto == 0 {
					fmt.Println("Buyed")
					crypto = float64(balance) / float64(int(mHistory[len(mHistory)-1].Close))
					balance = 0
				}
			}

			if isSell && !isBuy {
				if crypto != 0 {
					fmt.Println(mHistory[len(mHistory)-1].Close, price, "Selled")
					balance = int(crypto * float64(int(mHistory[len(mHistory)-1].Close)))
					crypto = 0
				}
			}

			if s%5 == 0 {
				LogDTO.Histories = append(LogDTO.Histories,
					[]string{
						fmt.Sprintf("isDown : %v , isDownEqual : %v , isUp : %v", isDown, isdownEqual, isUp),
					},
					[]string{
						fmt.Sprintf("isBuy   : %v, isSell   : %v", isBuy, isSell),
					},
					[]string{
						fmt.Sprintf("rest balance dummy:   : %v", balance),
					},
				)

				if crypto != 0 {
					LogDTO.Histories = append(LogDTO.Histories,
						[]string{
							fmt.Sprintf("BTC   : %v , Profit : %.2f", crypto, crypto*float64(int(mHistory[len(mHistory)-1].Close))),
						},
					)
				}
			}
		}

		if len(LogDTO.Histories) <= 0 {
			LogDTO.Histories = append(LogDTO.Histories, HeaderLog)
		}

	}
}

func (ig *IndodaxGUICtx) orderSell(price string) {
	mh := ig.getMarket()
	if len(mh) > 0 {
		o, _ := strconv.Atoi(price)
		if mh[len(mh)-1].Close >= int64(o) {
			dialog.NewError(errors.New("not in price buy"), ig.window).Show()
			return
		}
	}

	info, err := ig.exchange.Info(context.Background())
	if err != nil {
		dialog.NewError(err, ig.window).Show()
		return
	}
	var btc = ""
	if PairDTO == pairBTC {
		btc = info.Return.Balance.BTC
	} else if PairDTO == pairAlt {
		btc = info.Return.Balance.ALT
	} else if PairDTO == pairAbyss {
		btc = info.Return.Balance.ABYSS
	}

	res, err := ig.exchange.Order(context.Background(), repository.CreateOrderParams{
		OrderPrice:  price,
		Btc:         btc,
		OrderType:   "sell",
		OrderCrypto: PairDTO,
	})

	if err != nil {
		dialog.NewError(err, ig.window).Show()
		return
	}

	dialog.NewInformation(strconv.Itoa(int(res.Return.OrderID)), strconv.Itoa(int(res.Return.Fee)), ig.window).Show()
}

func (ig *IndodaxGUICtx) orderBuy(spend, price string) {
	mh := ig.getMarket()
	if len(mh) > 0 {
		o, _ := strconv.Atoi(price)
		if mh[len(mh)-1].Close <= int64(o) {
			dialog.NewError(errors.New("not in price buy"), ig.window).Show()
			return
		}
	}

	res, err := ig.exchange.Order(context.Background(), repository.CreateOrderParams{
		OrderPrice:  price,
		Idr:         spend,
		OrderType:   "buy",
		OrderCrypto: PairDTO,
	})
	if err != nil {
		dialog.NewError(err, ig.window).Show()
		return
	}

	dialog.NewInformation(strconv.Itoa(int(res.Return.OrderID)), strconv.Itoa(int(res.Return.Fee)), ig.window).Show()
}

func (ig *IndodaxGUICtx) getMarket() []models.MarketHistoryIndodax {
	payload := models.MarketHistoryPayload{
		Symbol:    SymbolDTO,
		TimeFrame: TimeFrimeDTO,
		From:      ig.config.Time.Now(nil).Add(time.Minute * -time.Duration(1440)).Unix(),
		To:        ig.config.Time.Now(nil).Unix(),
	}

	mHistory, err := ig.exchange.MarketHistory(context.Background(), payload)
	if err != nil {
		dialog := dialog.NewError(err, ig.window)
		dialog.Show()
		return nil
	}

	return mHistory
}

func (ig *IndodaxGUICtx) orderHistory() ([][]string, [][]string) {
	orderHistory := [][]string{HeaderOrderHistory}
	orderPendingHistory := [][]string{HeaderOrderHistory}
	result, err := ig.exchange.OrderHistory(context.Background(), PairDTO)
	if err != nil {
		dialog := dialog.NewError(err, ig.window)
		dialog.Show()
		return orderHistory, orderPendingHistory
	}

	sort.Slice(result.Return.Orders, func(i, j int) bool {
		orderOne, _ := strconv.ParseFloat(result.Return.Orders[i].FinishTime, 64)
		orderTwo, _ := strconv.ParseFloat(result.Return.Orders[j].FinishTime, 64)
		return orderOne > orderTwo
	})

	for _, history := range result.Return.Orders {
		if history.Status != "pending" {
			price, _ := strconv.ParseFloat(history.Price, 64)
			orderHistory = append(orderHistory, []string{
				history.OrderID,
				history.Type,
				rupiah.FormatRupiah(price),
				history.Status,
				ig.config.Time.TimpStampToDateStr(history.FinishTime, "02-01-2006 15:04"),
			})
		}
	}

	for _, history := range result.Return.Orders {
		if history.Status == "pending" {
			price, _ := strconv.ParseFloat(history.Price, 64)
			orderPendingHistory = append(orderPendingHistory, []string{
				history.Type,
				history.OrderID,
				rupiah.FormatRupiah(price),
				history.Status,
				ig.config.Time.TimpStampToDateStr(history.FinishTime, "02-01-2006 15:04"),
			})
		}
	}

	return orderHistory, orderPendingHistory
}

func (ig *IndodaxGUICtx) marketHistory() [][]string {
	marketHistory := [][]string{HeaderMarketHistory}

	mHistory := ig.getMarket()
	if len(mHistory) > 0 {
		for index := len(mHistory) - 1; index >= 0; index-- {
			marketHistory = append(marketHistory, []string{
				SymbolDTO,
				rupiah.FormatRupiah(float64(mHistory[index].High)),
				rupiah.FormatRupiah(float64(mHistory[index].Low)),
				rupiah.FormatRupiah(float64(mHistory[index].Open)),
				rupiah.FormatRupiah(float64(mHistory[index].Close)),
				ig.config.Time.TimpStampToDateStr(strconv.Itoa(int(mHistory[index].Time)), "02-01-2006 15:04"),
			})
		}
	}

	return marketHistory
}

func (ig *IndodaxGUICtx) marketHistoryV2() [][]string {
	marketHistory := [][]string{HeaderPrice}

	mHistory := ig.getMarket()
	if len(mHistory) > 0 {
		for index := len(mHistory) - 1; index >= 0; index-- {
			marketHistory = append(marketHistory, []string{
				rupiah.FormatRupiah(float64(mHistory[index].Close)),
			})
		}

		OlPrice = float64(mHistory[len(mHistory)-1].Close)
	}

	return marketHistory
}

func (ig *IndodaxGUICtx) marketPendingHistory() (sell [][]string, buy [][]string) {
	sell = [][]string{HeaderMarketPendingHistory}
	buy = [][]string{HeaderMarketPendingHistory}

	marketPending, err := ig.exchange.MarketPendingtHistory(context.Background(), SymbolPendingDTO)
	if err != nil {
		dialog := dialog.NewError(err, ig.window)
		dialog.Show()
		return
	}

	if SymbolPendingDTO == symbolPendingBTC {
		sort.Slice(marketPending.SellOrders, func(i, j int) bool {
			bidOne, _ := strconv.ParseFloat(marketPending.SellOrders[i].SumBTC, 64)
			bidTwo, _ := strconv.ParseFloat(marketPending.SellOrders[j].SumBTC, 64)
			return bidOne < bidTwo
		})
	} else if SymbolPendingDTO == symbolPendingETH {
		sort.Slice(marketPending.SellOrders, func(i, j int) bool {
			bidOne, _ := strconv.ParseFloat(marketPending.SellOrders[i].SumETH, 64)
			bidTwo, _ := strconv.ParseFloat(marketPending.SellOrders[j].SumETH, 64)
			return bidOne < bidTwo
		})
	}

	sort.Slice(marketPending.BuyOrders, func(i, j int) bool {
		priceOne, _ := strconv.ParseFloat(marketPending.BuyOrders[i].SumRP, 64)
		priceTwo, _ := strconv.ParseFloat(marketPending.BuyOrders[j].SumRP, 64)
		return priceOne < priceTwo
	})

	payload := models.MarketHistoryPayload{
		Symbol:    SymbolDTO,
		TimeFrame: TimeFrimeDTO,
		From:      ig.config.Time.Now(nil).Add(time.Minute * -time.Duration(1440)).Unix(),
		To:        ig.config.Time.Now(nil).Unix(),
	}

	mHistory, err := ig.exchange.MarketHistory(context.Background(), payload)
	if err != nil {
		dialog := dialog.NewError(err, ig.window)
		dialog.Show()
		return sell, buy
	}

	for index := len(marketPending.SellOrders) - 1; index >= 0; index-- {
		price, _ := strconv.ParseFloat(marketPending.SellOrders[index].Price, 64)
		payloadNumb := manipulator.PayloadNumber{}

		if SymbolPendingDTO == symbolPendingBTC {
			bid, _ := strconv.ParseFloat(marketPending.SellOrders[index].SumBTC, 64)
			payloadNumb.NumberOne = int(bid)

		} else if SymbolPendingDTO == symbolPendingETH {
			bid, _ := strconv.ParseFloat(marketPending.SellOrders[index].SumETH, 64)
			payloadNumb.NumberOne = int(bid)

		}
		sell = append(sell, []string{
			SymbolDTO,
			rupiah.FormatRupiah(payloadNumb.MakeNumberToBeFloat() * float64(mHistory[len(mHistory)-1].Close)),
			rupiah.FormatRupiah(price),
		})
	}

	for index := len(marketPending.BuyOrders) - 1; index >= 0; index-- {
		price, _ := strconv.ParseFloat(marketPending.BuyOrders[index].Price, 64)
		sumRp, _ := strconv.ParseFloat(marketPending.BuyOrders[index].SumRP, 64)

		buy = append(buy, []string{
			SymbolDTO,
			rupiah.FormatRupiah(sumRp),
			rupiah.FormatRupiah(price),
		})
	}

	return
}

func (ig *IndodaxGUICtx) chartIndocator() (times []time.Time, closePrices []float64) {
	mHistory := ig.getMarket()

	i := 0
	isExist := false
	for _, history := range mHistory {
		i++
		timeResult := ig.config.Time.TimpStampToDate(strconv.Itoa(int(history.Time)), "2006-01-02")
		for _, tm := range times {
			if timeResult == tm {
				isExist = true
			} else {
				isExist = false
			}
		}
		if !isExist {
			times = append(times, timeResult)
		}

		closePrices = append(closePrices, float64(history.Close))
		if i > 1000 {
			break
		}
	}

	return
}
