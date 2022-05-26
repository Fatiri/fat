package indodaxgui

import (
	"context"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fat/app/gui/component"
	"github.com/yudapc/go-rupiah"
)

var (
	HeaderPrice []string = []string{"Price Market"}
	MarketPrice          = &component.List{
		Histories: [][]string{HeaderLog},
	}
	bid     = widget.NewEntry()
	price   = widget.NewEntry()
	idr     = widget.NewEntry()
	crypto  = widget.NewEntry()
	profit  = widget.NewEntry()
	OlPrice = 0.00
)

func (ig *IndodaxGUICtx) ShowV2() *container.Split {
	logList := component.NewTable(ig.window).V5(MarketPrice)
	logList.OnSelected = func(id widget.TableCellID) {
		f1 := strings.ReplaceAll(MarketPrice.Histories[id.Row][id.Col], "Rp ", "")
		f2 := strings.ReplaceAll(f1, ".", "")
		price.SetText(f2)
	}
	go func(tbl *widget.Table) {
		tm := time.NewTicker(time.Second * 2)
		for range tm.C {
			MarketPrice.Histories = ig.marketHistoryV2()
			tbl.ScrollToTop()
		}
	}(logList)

	go func() {
		tm := time.NewTicker(time.Second * 2)
		for range tm.C {
			info, _ := ig.exchange.Info(context.Background())
			if info != nil {
				idr.SetText(rupiah.FormatRupiah(info.Return.Balance.IDR))
				crypto.SetText(info.Return.Balance.BTC)
				t2, _ := strconv.ParseFloat(info.Return.Balance.BTC, 64)
				if PairDTO == pairAlt {
					t2, _ = strconv.ParseFloat(info.Return.Balance.ALT, 64)
				}
				profit.SetText(rupiah.FormatRupiah(OlPrice * float64(t2)))
			}
		}
	}()
	return container.NewVSplit(
		container.NewGridWrap(fyne.NewSize(500, 200),
			container.NewAppTabs(
				container.NewTabItem("Buy", container.NewCenter(container.NewGridWrap(fyne.NewSize(400, 100), ig.formBuy()))),
				container.NewTabItem("Sell", container.NewCenter(container.NewGridWrap(fyne.NewSize(400, 100), ig.formSell()))),
				container.NewTabItem("Info", container.NewCenter(container.NewGridWrap(fyne.NewSize(400, 100), ig.formInfo()))),
			),
		),
		container.NewGridWrap(fyne.NewSize(500, 600), logList),
	)
}

func (ig *IndodaxGUICtx) formBuy() *widget.Form {
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "BID", Widget: bid},
			{Text: "Price Market", Widget: price}},
		OnSubmit: func() {
			if bid.Text != "" {
				ig.orderBuy(bid.Text, price.Text)
			}
		},
	}

	form.SubmitText = "BUY"

	return form
}

func (ig *IndodaxGUICtx) formInfo() *widget.Form {
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Balance", Widget: idr},
			{Text: "crypto", Widget: crypto},
			{Text: "Profit", Widget: profit}},
	}

	crypto.Disable()
	idr.Disable()
	profit.Disable()

	return form
}

func (ig *IndodaxGUICtx) formSell() *widget.Form {
	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Price Market", Widget: price}},
		OnSubmit: func() {
			if price.Text != "" {
				ig.orderSell(price.Text)
			}
		},
	}

	form.SubmitText = "SELL"

	return form
}
