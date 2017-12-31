package k

import "github.com/jroimartin/gocui"

// ModalView is the UI state of the modal component
type ModalView struct {
	Title    string
	Kind     ModalKind
	Cursor   int
	Selected string
	Lines    []string
	Size     ModalSize
}

const (
	// KindDebug is a special debug modal
	KindDebug ModalKind = "Debug"
	// KindModalNamespaces represents a modal of the type namespace
	KindModalNamespaces ModalKind = "Namespaces"
	// KindModalResources represents a modal of the type resource
	KindModalResources ModalKind = "Resources"
	// KindModalSelectContainer represents a modal of the type container selection
	KindModalSelectContainer ModalKind = "Container"
	// KindModalContainerLogs represents a modal of the type container logs
	KindModalContainerLogs ModalKind = "ContainerLogs"
)

// ModalSize is the size of the modal
type ModalSize string

const (

	// ModalSizeSmall represents a Small modal
	ModalSizeSmall = "Small"
	// ModalSizeMedium represents a Medium modal
	ModalSizeMedium = "Medium"
	// ModalSizeLarge represents a Large modal
	ModalSizeLarge = "Large"
	// ModalSizeExtraLarge represents a ExtraLarge modal
	ModalSizeExtraLarge = "ExtraLarge"
)

//SetSize updates the modal UI state's size
func (mv *ModalView) SetSize(g1 *gocui.Gui, size ModalSize) {
	Debugln("ModalView: SetSize")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Size = size
			return nil
		},
	)
}

//SetKind updates the modal UI state's kind
func (mv *ModalView) SetKind(g1 *gocui.Gui, kind ModalKind) {
	Debugln("ModalView: SetKind")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Kind = kind
			return nil
		},
	)
}

//SetLines updates the modal UI state's lines
func (mv *ModalView) SetLines(g1 *gocui.Gui, lines []string) {
	Debugln("ModalView: SetLines")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Lines = lines
			return nil
		},
	)
}

//SetTitle updates the modal UI state's title
func (mv *ModalView) SetTitle(g1 *gocui.Gui, title string) {
	Debugln("ModalView: SetTitle")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Title = title
			return nil
		},
	)
}

//SetCursor updates the modal UI state's cursor position
func (mv *ModalView) SetCursor(g1 *gocui.Gui, pos int) {
	Debugln("ModalView: SetCursor")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Cursor = pos
			if len(mv.Lines) > 0 && mv.Cursor > 0 {
				mv.Selected = mv.Lines[mv.Cursor-1]
			}
			mv.Selected = ""
			return nil
		},
	)
}

// CursorMove updates the UI state with a new cursor position (delta)
func (mv *ModalView) CursorMove(g1 *gocui.Gui, delta int) {
	Debugln("ModalView: CursorMove")
	g1.Update(
		func(g *gocui.Gui) error {
			mv.Cursor = mv.Cursor + delta
			switch {
			case mv.Cursor < 0:
				mv.Cursor = 0
			case mv.Cursor > len(mv.Lines)-1:
				mv.Cursor = len(mv.Lines) - 1
			}
			return nil
		},
	)
}
