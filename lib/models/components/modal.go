package components

import (
	"github.com/hachi-n/everlasting/lib/pkgerutil"
	"github.com/rivo/tview"
)

type Modal struct {
	app *tview.Application
	primitive *tview.Modal
	executeFilepath string
}


var (
	modalComponent *Modal
)

func NewModal(app *tview.Application, executeFilepath string ) *Modal {


	modal := tview.NewModal().
		SetText("Do you want to quit the application?").
		AddButtons([]string{"OK", "Quit", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			} else if buttonLabel == "Cancel" {

			} else if buttonLabel == "OK" {
				app.Stop()
				pkgerutil.Exec(executeFilepath)
			}
		})

	modalComponent = new(Modal)
	modalComponent.app = app
	modalComponent.primitive = modal
	modalComponent.executeFilepath = executeFilepath

	return modalComponent

}

func (p *Modal) Primitive() tview.Primitive {
	p.primitive = modalComponent.primitive
	return modalComponent.primitive
}

func (p *Modal) Display() {
	p.app.SetRoot(p.primitive, true)
}
