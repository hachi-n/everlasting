package everlasting

import (
	"github.com/hachi-n/everlasting/lib/models"
)

func TuiStart(resourcesDir string, packedFlag bool) error {
	tui := models.NewTui(resourcesDir)
	return tui.Display()
}
