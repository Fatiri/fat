package indodaxgui

import (
	"fmt"
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/fat/app/gui/component"
	"github.com/fat/models"
	"github.com/fat/usecase/exchange"
	"github.com/fyne-io/examples/img/icon"
)

type IndodaxGUI interface {
	Show() *container.AppTabs
	ShowV2() *container.Split
}

type IndodaxGUICtx struct {
	config   *models.Config
	exchange exchange.Indodax
	window   fyne.Window
}

func NewIndodaxGUI(window fyne.Window, conf *models.Config) IndodaxGUI {
	return &IndodaxGUICtx{
		window:   window,
		config:   conf,
		exchange: exchange.NewIndodax(conf),
	}
}

const (
	pairBTC          = "btc_idr"
	pairAlt          = "alt_idr"
	pairAbyss        = "abyss_idr"
	symbolBTC        = "btcidr"
	symbolAbyss      = "abyssidr"
	symbolALT        = "altidr"
	symbolPendingBTC = "BTCIDR"
	pairETH          = "eth_idr"
	symbolETH        = "ethidr"
	symbolPendingETH = "ETHIDR"
)

type IndikatorsBot struct {
	RSI    bool
	MACD   bool
	BBANDS bool
}

type Indikators struct {
	RSI    []float64
	MACD   []float64
	BBANDS []float64
}

type StatusBot struct {
	Run          bool
	OrderSpendRp float64
	Profit       float64
}

type TradesTrategy struct {
	Buy      bool
	Sell     bool
	RSIBuy   bool
	RSISell  bool
	MACDBuy  bool
	MACDSell bool
}

var (
	HeaderOrderHistory         []string = []string{"Order ID", "Type", "Price Market", "Status", "Time"}
	HeaderMarketHistory        []string = []string{"Crypto", "Price High", "Price Low", "Price Open", "Price Close", "Time Close"}
	HeaderMarketPendingHistory []string = []string{"Crypto", "BID", "Price BID"}
	HeaderLog                  []string = []string{"Info"}
	OrderHistoryDTO                     = &component.List{
		Histories: [][]string{HeaderOrderHistory},
	}
	OrderPendingHistoryDTO = &component.List{
		Histories: [][]string{HeaderOrderHistory},
	}
	MarketHistoryDTO = &component.List{
		Histories: [][]string{HeaderMarketHistory},
	}

	MarketPendingSellDTO = &component.List{
		Histories: [][]string{HeaderMarketPendingHistory},
	}
	MarketPendingBuyDTO = &component.List{
		Histories: [][]string{HeaderMarketPendingHistory},
	}
	LogDTO = &component.List{
		Histories: [][]string{HeaderLog},
	}
	PairDTO          string = pairAlt
	TimeFrimeDTO     string = "1"
	SymbolDTO        string = symbolALT
	SymbolPendingDTO string = ""
	IndikatorsBotDTO        = IndikatorsBot{}
	IndikatorsDTO           = Indikators{}
	BotStatusDTO            = &StatusBot{}
	TradesTrategyDTO        = &TradesTrategy{}
)

func (ig *IndodaxGUICtx) Show() *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItem("Trade", ig.trade()),
		container.NewTabItem("Profile", widget.NewLabel("Profile")),
	)

	return tabs
}

func (ig *IndodaxGUICtx) trade() *container.Split {
	content := container.NewMax()

	appList := ig.appList(content)

	// size windows
	form := container.NewHSplit(
		container.NewGridWrap(
			fyne.NewSize(215, 900),
			container.New(layout.NewMaxLayout(), appList),
		),
		container.NewGridWrap(
			fyne.NewSize(1570, 900),
			container.NewGridWrap(
				fyne.NewSize(1565, 900),
				content,
			),
		),
	)

	return form
}

func (ig *IndodaxGUICtx) appList(content *fyne.Container) *widget.List {
	itemList := component.NewItemList(ig.window)
	var apps = []component.AppInfo{}

	apps = append(apps,
		component.AppInfo{Name: "BTC", Icon: icon.XKCDBitmap, Canv: false, Run: ig.content},
		component.AppInfo{Name: "ETH", Icon: icon.XKCDBitmap, Canv: false, Run: ig.content},
	)

	appList := itemList.V1(apps, content)

	appList.OnSelected = func(id widget.ListItemID) {
		if id == 0 {
			PairDTO = pairBTC
			SymbolDTO = symbolBTC
			SymbolPendingDTO = symbolPendingBTC
		} else if id == 1 {
			PairDTO = pairETH
			SymbolDTO = symbolETH
			SymbolPendingDTO = symbolPendingETH
		}
		content.Objects = []fyne.CanvasObject{apps[id].Run(ig.window)}
	}

	return appList
}

func (ig *IndodaxGUICtx) content(win fyne.Window) fyne.CanvasObject {
	form := container.NewVSplit(
		container.NewGridWrap(
			fyne.NewSize(1560, 280),
			container.NewHSplit(
				container.NewGridWrap(
					fyne.NewSize(580, 280),
					ig.tabOrderTabs(),
				),
				container.NewGridWrap(
					fyne.NewSize(975, 280),
					ig.tabOrderHistoryTabs(),
				),
			),
		),
		container.NewGridWrap(
			fyne.NewSize(1565, 630),
			container.NewMax(
				ig.tabMarketTabs(),
			),
		),
	)
	return form
}

func (ig *IndodaxGUICtx) tabOrderTabs() *container.AppTabs {
	orderManual := component.NewForm(ig.window).V1()
	orderManual.SubmitText = "Order"
	orderManual.OnSubmit = func() {
		log.Println("Form submitted:")

	}

	profit := component.NewForm(ig.window).V2()

	fmt.Println()
	return container.NewAppTabs(
		container.NewTabItem("Trade Manual", container.NewCenter(container.NewGridWrap(fyne.NewSize(550, 180), orderManual))),
		container.NewTabItem("Trade Automatic", container.NewGridWrap(fyne.NewSize(580, 180), ig.tradeAutomaticTabs())),
		container.NewTabItem("Calculation Profit", container.NewCenter(container.NewGridWrap(fyne.NewSize(550, 180), profit))),
	)
}

func (ig *IndodaxGUICtx) tradeAutomaticTabs() *container.AppTabs {
	bot := component.NewForm(ig.window).V3(BotStatusDTO.Run)
	bot.SubmitText = "Run"
	bot.OnSubmit = func() {
		if BotStatusDTO.Run {
			dg := dialog.NewInformation("Bot Infomation", "Bot currently running", ig.window)
			dg.Show()
			return
		}
		BotStatusDTO.Run = true
	}
	bot.CancelText = "Stop"
	bot.OnCancel = func() {
		if !BotStatusDTO.Run {
			dg := dialog.NewInformation("Bot Infomation", "Bot currently stoped", ig.window)
			dg.Show()
			return
		}
		BotStatusDTO.Run = false
	}

	indicators, listCheck := component.NewForm(ig.window).V4()

	listCheck[2].OnChanged = func(b bool) {
		IndikatorsBotDTO = IndikatorsBot{
			RSI:    listCheck[0].Checked,
			MACD:   listCheck[1].Checked,
			BBANDS: listCheck[2].Checked,
		}
	}

	logList := component.NewTable(ig.window).V4(LogDTO)
	go ig.BOT()
	go func(tbl *widget.Table) {
		tm := time.NewTicker(time.Second * 5)
		for range tm.C {
			tbl.ScrollToBottom()
		}
	}(logList)

	return container.NewAppTabs(
		container.NewTabItem("Bot Trade", container.NewCenter(container.NewGridWrap(fyne.NewSize(550, 130), bot))),
		container.NewTabItem("Indikators", container.NewCenter(container.NewGridWrap(fyne.NewSize(400, 130), indicators))),
		container.NewTabItem("Logs", container.NewGridWrap(fyne.NewSize(580, 190), logList)),
	)
}

func (ig *IndodaxGUICtx) tabOrderHistoryTabs() *container.AppTabs {
	tableOH := component.NewTable(ig.window).V1(OrderHistoryDTO)
	tableOPH := component.NewTable(ig.window).V1(OrderPendingHistoryDTO)

	ct := container.NewAppTabs(
		container.NewTabItem("My Pending Order", tableOPH),
		container.NewTabItem("My Order History", tableOH),
	)

	orderHistory, orderPendingHistory := ig.orderHistory()
	OrderHistoryDTO.Histories = orderHistory
	OrderPendingHistoryDTO.Histories = orderPendingHistory

	go func(tblOH *widget.Table, tblOPH *widget.Table) {
		tc := time.NewTicker(time.Second * time.Duration(60))
		for range tc.C {
			orderHistory, orderPendingHistory := ig.orderHistory()
			OrderHistoryDTO.Histories = orderHistory
			OrderPendingHistoryDTO.Histories = orderPendingHistory
			tblOH.ScrollToTop()
			tblOPH.ScrollToTop()
		}
	}(tableOH, tableOPH)

	tableOPH.OnSelected = func(id widget.TableCellID) {
		fmt.Println(id)
	}
	return ct
}

func (ig *IndodaxGUICtx) tabMarketTabs() *container.AppTabs {
	tableMS := component.NewTable(ig.window).V2(MarketHistoryDTO)
	MarketHistoryDTO.Histories = ig.marketHistory()

	go func(tbl *widget.Table) {
		tc := time.NewTicker(time.Second * time.Duration(60))
		for range tc.C {
			MarketHistoryDTO.Histories = ig.marketHistory()
			tbl.ScrollToTop()
		}

	}(tableMS)

	return container.NewAppTabs(
		container.NewTabItem("Market History", tableMS),
		container.NewTabItem("Market Pending Order", ig.marketPendingOrderContainer()),
		container.NewTabItem("Graph", component.Chart(ig.chartIndocator())),
	)
}

func (ig *IndodaxGUICtx) marketPendingOrderContainer() *fyne.Container {
	tablePOS := component.NewTable(ig.window).V3(MarketPendingSellDTO)
	tablePOB := component.NewTable(ig.window).V3(MarketPendingBuyDTO)
	MarketPendingSellDTO.Histories, MarketPendingBuyDTO.Histories = ig.marketPendingHistory()

	go func(tblPOS *widget.Table, tblPOB *widget.Table) {
		tc := time.NewTicker(time.Second * time.Duration(60))
		for range tc.C {
			MarketPendingSellDTO.Histories, MarketPendingBuyDTO.Histories = ig.marketPendingHistory()
			tblPOS.ScrollToTop()
			tblPOB.ScrollToTop()
		}

	}(tablePOS, tablePOB)

	return container.NewGridWrap(
		fyne.NewSize(1560, 710),
		container.NewHSplit(
			container.NewGridWrap(
				fyne.NewSize(775, 710),
				container.NewAppTabs(
					container.NewTabItem("Sell", tablePOS),
				),
			),
			container.NewGridWrap(
				fyne.NewSize(775, 710),
				container.NewAppTabs(
					container.NewTabItem("Buy", tablePOB),
				),
			),
		),
	)
}
