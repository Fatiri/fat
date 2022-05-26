package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Table interface {
	V1(data *List) *widget.Table
	V2(data *List) *widget.Table
	V3(data *List) *widget.Table
	V4(data *List) *widget.Table
	V5(data *List) *widget.Table
}

type TableCtx struct {
	window fyne.Window
}

func NewTable(window fyne.Window) Table {
	return &TableCtx{
		window: window,
	}
}

func (t *TableCtx) V1(data *List) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data.Histories), len(data.Histories[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			o.(*widget.Label).SetText(data.Histories[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 192)
	list.SetColumnWidth(1, 192)
	list.SetColumnWidth(2, 192)
	list.SetColumnWidth(3, 192)
	list.SetColumnWidth(4, 192)
	return list
}

func (t *TableCtx) V2(data *List) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data.Histories), len(data.Histories[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			o.(*widget.Label).SetText(data.Histories[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 260)
	list.SetColumnWidth(1, 260)
	list.SetColumnWidth(2, 260)
	list.SetColumnWidth(3, 260)
	list.SetColumnWidth(4, 260)
	list.SetColumnWidth(5, 260)
	return list
}

func (t *TableCtx) V3(data *List) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data.Histories), len(data.Histories[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			o.(*widget.Label).SetText(data.Histories[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 256)
	list.SetColumnWidth(1, 256)
	list.SetColumnWidth(2, 256)
	return list
}

func (t *TableCtx) V4(data *List) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data.Histories), len(data.Histories[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			o.(*widget.Label).SetText(data.Histories[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 580)

	return list
}

func (t *TableCtx) V5(data *List) *widget.Table {
	list := widget.NewTable(
		func() (int, int) {
			return len(data.Histories), len(data.Histories[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).Alignment = fyne.TextAlignCenter
			o.(*widget.Label).TextStyle = fyne.TextStyle{
				Bold:   true,
				Italic: true,
			}
			o.(*widget.Label).SetText(data.Histories[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 500)

	return list
}
