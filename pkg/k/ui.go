package k

import (
	"github.com/jroimartin/gocui"
)

// UI Reducer is a high level reducer that contains state of the different UI components in the app
type UIReducer struct {
	Table        *TableView
	Modal        *ModalView
	State        *StateView
	Debug        *DebugView
	ActiveScreen Screen
}

// DebugView is the state for the Debug component
type DebugView struct {
	Cursor int
}

// StateView is the state for the State component
type StateView struct {
	Cursor int
}

// CursorMove updates the UI state with a new cursor position (delta)
func (ur *UIReducer) CursorMove(g1 *gocui.Gui, delta int) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.Modal.Cursor = ur.Modal.Cursor + delta
			switch {
			case ur.Modal.Cursor < 0:
				ur.Modal.Cursor = 0
			case ur.Modal.Cursor > len(ur.Modal.Lines)-1:
				ur.Modal.Cursor = len(ur.Modal.Lines) - 1
			default:
				Debugln("CursorMove: Unsupported Screen", ur.ActiveScreen)
				return nil
			}
			if len(ur.Modal.Lines) > 0 {
				ur.Modal.Selected = ur.Modal.Lines[ur.Modal.Cursor]
			}

			return nil
		},
	)
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
