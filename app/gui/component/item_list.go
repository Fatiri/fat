package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type ItemList interface {
	V1(apps []AppInfo, content *fyne.Container) *widget.List
	V2(items *[]string) *widget.List
}

type ItemListCtx struct {
	window fyne.Window
}

func NewItemList(window fyne.Window) ItemList {
	return &ItemListCtx{
		window: window,
	}
}

func (il *ItemListCtx) V1(apps []AppInfo, content *fyne.Container) *widget.List {

	appList := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height
			icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
			return container.NewBorder(nil, nil, icon, nil,
				label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = apps[id].Icon
			img.Refresh()
			text.SetText(apps[id].Name)
		})

	return appList
}

func (il *ItemListCtx) V2(items *[]string) *widget.List {
	data := binding.BindStringList(
		items,
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	return list
}
