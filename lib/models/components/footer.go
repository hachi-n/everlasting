package components

import "github.com/rivo/tview"

type Footer struct {
	app       *tview.Application
	primitive *tview.TextView
}

var (
	footerComponent *Footer
)

func NewFooter(app *tview.Application) *Footer {
	if footerComponent != nil {
		return footerComponent
	}

	newPrimitive := func(text string) *tview.TextView {
		textView := tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
		return textView
	}

	footerMessage := newPrimitive("footer")

	footerComponent = new(Footer)
	footerComponent.app = app
	footerComponent.primitive = footerMessage

	return footerComponent
}

func (p *Footer) Primitive() tview.Primitive {
	p.primitive = footerComponent.primitive
	return footerComponent.primitive
}
