package components

import (
	"bytes"
	"github.com/markbates/pkger"
	"github.com/rivo/tview"
)

type Menu struct {
	app       *tview.Application
	primitive *tview.TextView
}

var (
	menuComponent *Menu
)

const introduction = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.

Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.

Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.

[yellow]Press Enter, then Tab/Backtab for word selections`

func NewMenu(app *tview.Application) *Menu {
	if menuComponent != nil {
		return menuComponent
	}

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true).
		SetText(introduction).
		SetChangedFunc(func() {
			app.Draw()
		})

	menuComponent = new(Menu)
	menuComponent.app = app
	menuComponent.primitive = textView

	return menuComponent
}

func (p *Menu) Primitive() tview.Primitive {
	p.primitive = menuComponent.primitive
	return menuComponent.primitive
}

func (p *Menu) updateViewFromFile(filePath string) {
	f, err := pkger.Open(filePath)
	if err == nil {
		buf := bytes.Buffer{}
		_, err:= buf.ReadFrom(f)
		if err == nil {
			p.primitive.SetText(buf.String())
		}
	}
}

func (p *Menu) updateViewFromString(message string) {
	p.primitive.SetText(introduction)
}

