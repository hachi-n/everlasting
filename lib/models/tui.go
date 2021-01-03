package models

import (
	"github.com/gdamore/tcell/v2"
	"github.com/hachi-n/everlasting/lib/models/components"
	"github.com/rivo/tview"
)

type Tui struct {
	app    *tview.Application
	layout *Layout
}

const (
	title = "everlasting"
)

func NewTui(resoucesDir string) *Tui {
	//tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault
	tview.Styles.PrimitiveBackgroundColor = tcell.Color235
	tview.Styles.ContrastBackgroundColor = tcell.Color235

	tui := new(Tui)
	tui.app = tview.NewApplication()

	layout := NewLayout(
		tui.app,
		components.NewSideNav(tui.app, resoucesDir),
		components.NewMenu(tui.app),
		components.NewHeader(tui.app),
		components.NewFooter(tui.app),
	)

	tui.layout = layout
	return tui
}

func (t *Tui) Display() error {
	return t.app.SetRoot(t.layout.Primitive(), true).Run()
}

func (t *Tui) Stop() {
	t.app.Stop()
}
