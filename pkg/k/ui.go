package k

import (
	"github.com/jroimartin/gocui"
)

// UIReducer is a high level reducer that contains state of the different UI components in the app
type UIReducer struct {
	Table        *TableView
	Modal        *ModalView
	ActiveScreen Screen
}

// SetActiveScreen updates the UI state with a new active screen
func (ur *UIReducer) SetActiveScreen(g1 *gocui.Gui, screen Screen) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = screen
			return nil
		},
	)
}
