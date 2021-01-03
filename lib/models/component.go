package models

import "github.com/rivo/tview"

type Component interface {
	Primitive() tview.Primitive
}
