package component

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Form interface {
	V1() *widget.Form
	V2() *widget.Form
	V3(botRun bool) *widget.Form
	V4() (form *widget.Form, checkGroups []*widget.Check)
	Login() (form *widget.Form, token *widget.Entry)
}

type FormCtx struct {
	window fyne.Window
}

func NewForm(window fyne.Window) Form {
	return &FormCtx{
		window: window,
	}
}

func (f *FormCtx) V1() *widget.Form {
	price := widget.NewEntry()
	bid := widget.NewEntry()
	lose := widget.NewEntry()

	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Market Price", Widget: price},
			{Text: "BID", Widget: bid},
			{Text: "Lose Price", Widget: lose}},
	}
}

func (f *FormCtx) V2() *widget.Form {
	Balance := widget.NewEntry()
	Profit := widget.NewEntry()
	Balance.Disable()
	Profit.Disable()

	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Spend Money", Widget: Balance},
			{Text: "Profit %", Widget: Profit}},
	}
}

func (f *FormCtx) V3(botRun bool) *widget.Form {
	tradeRunning := widget.NewSelectEntry([]string{"1", "2", "3", "4", "5"})

	tradeRunning.OnChanged = func(s string) {
		fmt.Println(s)
	}

	bid := widget.NewEntry()
	lose := widget.NewEntry()

	return &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Do Trade  x", Widget: tradeRunning},
			{Text: "BID", Widget: bid},
			{Text: "Lose Price", Widget: lose}},
	}
}

func (f *FormCtx) V4() (form *widget.Form, indikators []*widget.Check) {
	RSI := widget.NewCheck("RSI", func(b bool) {})
	MACD := widget.NewCheck("MACD", func(b bool) {})
	BBANDS := widget.NewCheck("BBANDS", func(b bool) {})
	RSI.Checked = true
	MACD.Checked = true
	RSI.Disable()
	MACD.Disable()

	indikators = append(indikators, RSI, MACD, BBANDS)

	form = &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Indikators", Widget: RSI},
			{Text: "", Widget: MACD},
			{Text: "", Widget: BBANDS},
		},
	}
	return
}

func (f *FormCtx) Login() (form *widget.Form, token *widget.Entry) {
	token = widget.NewEntry()

	form = &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Token", Widget: token}},
	}

	return form, token
}