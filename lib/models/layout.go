package models

import (
	"github.com/rivo/tview"
)

type Layout struct {
	app       *tview.Application
	primitive tview.Primitive
}

func NewLayout(app *tview.Application, sidenav, menu, header, footer Component) *Layout {
	grid := tview.NewGrid().
		SetRows(2, 0, 2).
		SetColumns(30, 0).
		SetBorders(true).
		AddItem(header.Primitive(), 0, 0, 1, 3, 0, 0, false).
		AddItem(footer.Primitive(), 2, 0, 1, 3, 0, 0, false).
		AddItem(sidenav.Primitive(), 1, 0, 1, 1, 0, 100, true).
		AddItem(menu.Primitive(), 1, 1, 1, 2, 0, 100, false)
	layout := new(Layout)
	layout.app = app
	layout.primitive = grid

	return layout
}

func (p *Layout) Primitive() tview.Primitive {
	return p.primitive
}
