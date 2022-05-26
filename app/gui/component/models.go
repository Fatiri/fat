package component

import "fyne.io/fyne/v2"

type List struct {
	Histories [][]string
}

type AppInfo struct {
	Name string
	Icon fyne.Resource
	Canv bool
	Run  func(fyne.Window) fyne.CanvasObject
}

type ComponentBot struct {
	Run bool
}
