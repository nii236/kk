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
func (p *ModalView) SetSize(g1 *gocui.Gui, size ModalSize) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Size = size
			return nil
		},
	)
}

//SetKind updates the modal UI state's kind
func (p *ModalView) SetKind(g1 *gocui.Gui, kind ModalKind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

//SetLines updates the modal UI state's lines
func (p *ModalView) SetLines(g1 *gocui.Gui, lines []string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

//SetTitle updates the modal UI state's title
func (p *ModalView) SetTitle(g1 *gocui.Gui, title string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Title = title
			return nil
		},
	)
}

//SetCursor updates the modal UI state's cursor position
func (p *ModalView) SetCursor(g1 *gocui.Gui, pos int) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor = pos
			if len(p.Lines) > 0 && p.Cursor > 0 {
				p.Selected = p.Lines[p.Cursor-1]
			}
			p.Selected = ""
			return nil
		},
	)
}

//SetSelected updates the modal UI state's selection
func (p *ModalView) SetSelected(g1 *gocui.Gui, selected string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Selected = selected
			return nil
		},
	)
}
