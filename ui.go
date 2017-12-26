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

type TableView struct {
	Cursor   int
	Selected string
	Kind     Kind
	Lines    [][]string
}

type ModalView struct {
	Title    string
	Cursor   int
	Selected string
	Lines    []string
}

func (p *DebugView) CursorUp(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor--
			if p.Cursor < 0 {
				p.Cursor = 0
			}
			return nil
		},
	)
}

func (p *DebugView) CursorDown(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor++
			return nil
		},
	)
}

func (p *StateView) CursorUp(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor--
			if p.Cursor < 0 {
				p.Cursor = 0
			}
			return nil
		},
	)
}

func (p *StateView) CursorDown(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor++
			return nil
		},
	)
}

func (p *ModalView) CursorUp(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			if len(p.Lines) < 1 {
				p.Cursor = 0
				return nil
			}
			p.Cursor--
			if p.Cursor < 0 {
				p.Cursor = 0
			}
			p.Selected = p.Lines[p.Cursor]
			return nil
		},
	)
}

func (p *ModalView) CursorDown(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			if len(p.Lines) < 1 {
				p.Cursor = 0
				return nil
			}
			p.Cursor++
			if p.Cursor >= len(p.Lines) {
				p.Cursor = len(p.Lines) - 1
			}
			p.Selected = p.Lines[p.Cursor]
			return nil
		},
	)
}

func (p *TableView) CursorUp(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			if len(p.Lines) < 1 {
				p.Cursor = 0
				return nil
			}
			p.Cursor--
			if p.Cursor < 0 {
				p.Cursor = 0
			}
			p.Selected = p.Lines[p.Cursor][0]
			return nil
		},
	)
}

func (p *TableView) CursorDown(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			if len(p.Lines) < 1 {
				p.Cursor = 0
				return nil
			}
			p.Cursor++
			if p.Cursor >= len(p.Lines) {
				p.Cursor = len(p.Lines) - 1
			}
			p.Selected = p.Lines[p.Cursor][0]
			return nil
		},
	)
}
func (ur *UIReducer) SetTableActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenTable
			return nil
		},
	)

}
func (ur *UIReducer) SetStateActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenState
			return nil
		},
	)

}
func (ur *UIReducer) SetDebugActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenDebug
			return nil
		},
	)
}
func (ur *UIReducer) SetModalActive(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			ur.ActiveScreen = ScreenModal
			return nil
		},
	)
}

func (p *TableView) SetTableKind(g1 *gocui.Gui, kind Kind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

func (p *TableView) SetTableLines(g1 *gocui.Gui, lines [][]string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

func (p *ModalView) SetModalLines(g1 *gocui.Gui, lines []string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

func (p *ModalView) SetModalTitle(g1 *gocui.Gui, title string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Title = title
			return nil
		},
	)
}
