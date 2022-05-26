package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/fat/app/gui/component"
	indodaxgui "github.com/fat/app/gui/indodax"
	"github.com/fat/middleware"
	"github.com/fat/models"
	"github.com/fat/repository"
	"github.com/google/uuid"
)

type GUI interface {
	Run()
	RunV2()
}

type GUICtx struct {
	config *models.Config
}

func NewGUI(conf *models.Config) GUI {
	return &GUICtx{
		config: conf,
	}
}

func (g *GUICtx) Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("FAT")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(500, 100))
	myWindow.FullScreen()
	indodax := indodaxgui.NewIndodaxGUI(myWindow, g.config)

	tabs := container.NewAppTabs(
		container.NewTabItem("Indodax", indodax.Show()),
		container.NewTabItem("Toko Crypto", indodax.Show()),
		container.NewTabItem("Binance", indodax.Show()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	formLogin, token := component.NewForm(myWindow).Login()
	formLogin.OnSubmit = func() {

		err := g.checkLogin(token.Text)
		if err != nil {
			dialog.NewError(err, myWindow).Show()
		} else {
			formLogin.Hide()
			myWindow.SetContent(tabs)
			myWindow.Resize(fyne.NewSize(1700, 1000))
			myWindow.CenterOnScreen()
		}
	}

	myWindow.SetContent(formLogin)
	myWindow.ShowAndRun()
}

func (g *GUICtx) checkLogin(token string) error {
	auth, _ := middleware.NewAuthentication(g.config)
	_, err := auth.VerifyToken(token)

	str, _ := auth.CreateToken(repository.Account{
		AccountID:   uuid.New(),
		Username:    "ilham",
		AccountType: "ADMIN",
	})

	fmt.Println(str)

	return err
}

func (g *GUICtx) RunV2() {
	myApp := app.New()
	myWindow := myApp.NewWindow("FAT")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(500, 750))
	indodax := indodaxgui.NewIndodaxGUI(myWindow, g.config)

	tabs := container.NewAppTabs(
		container.NewTabItem("Indodax", indodax.ShowV2()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
