package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/kk"
)

// HandleDebugEsc runs when pressing esc while focused on a Modal
func HandleDebugEsc(s *k.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Modal: Pressed esc")
		s.UI.SetActiveScreen(g, k.ScreenTable)

		return nil
	}
}
