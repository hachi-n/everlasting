package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Header struct {
	app       *tview.Application
	primitive *tview.TextView
}

var (
	headerComponent *Header
)

func NewHeader(app *tview.Application) *Header {
	if headerComponent != nil {
		return headerComponent
	}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		switch key {
		case tcell.KeyESC:
			app.Stop()
		case tcell.KeyF1:
			app.SetFocus(sideNavComponent.primitive)
		default:
			return event
		}
		return nil
	})

	button := tview.NewTextView().SetText("ESC Quit F1 Quit")

	headerComponent = new(Header)
	headerComponent.app = app
	headerComponent.primitive = button

	return headerComponent
}

func (p *Header) Primitive() tview.Primitive {
	p.primitive = headerComponent.primitive
	return headerComponent.primitive
}

func (p *Header) Active() {

}
