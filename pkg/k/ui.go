package k

import (
	"github.com/jroimartin/gocui"
)

type UIReducer struct {
	Table        *TableView
	Modal        *ModalView
	State        *StateView
	Debug        *DebugView
	ActiveScreen Screen
}

type DebugView struct {
	Cursor int
}

type StateView struct {
	Cursor int
}

func (ur *UIReducer) CursorMove(g1 *gocui.Gui, delta int) {
	g1.Update(
		func(g *gocui.Gui) error {
			switch ur.ActiveScreen {
			case ScreenDebug:
				ur.Debug.Cursor = ur.Debug.Cursor + delta
				if ur.Debug.Cursor < 0 {
					ur.Debug.Cursor = 0
				}
			case ScreenModal:
				ur.Modal.Cursor = ur.Modal.Cursor + delta
				if ur.Modal.Cursor < 0 {
					ur.Modal.Cursor = 0
				}
				if ur.Modal.Cursor > len(ur.Modal.Lines)-1 {
					ur.Modal.Cursor = len(ur.Modal.Lines) - 1
				}
				if len(ur.Modal.Lines) > 0 {
					ur.Modal.Selected = ur.Modal.Lines[ur.Modal.Cursor]
				}

			}
			return nil
		},
	)
}

func (ur *UIReducer) SetActiveScreen(g1 *gocui.Gui, screen Screen) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = screen
			return nil
		},
	)
}
