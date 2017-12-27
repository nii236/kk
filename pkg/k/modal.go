package k

import "github.com/jroimartin/gocui"

type ModalView struct {
	Title    string
	Kind     ModalKind
	Cursor   int
	Selected string
	Lines    []string
	Size     ModalSize
}

const (
	KindModalNamespaces      ModalKind = "Namespaces"
	KindModalResources       ModalKind = "Resources"
	KindModalSelectContainer ModalKind = "Container"
	KindModalContainerLogs   ModalKind = "ContainerLogs"
)

type ModalSize string

const (
	ModalSizeSmall      = "Small"
	ModalSizeMedium     = "Medium"
	ModalSizeLarge      = "Large"
	ModalSizeExtraLarge = "ExtraLarge"
)

func (p *ModalView) SetSize(g1 *gocui.Gui, size ModalSize) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Size = size
			return nil
		},
	)
}

func (p *ModalView) SetKind(g1 *gocui.Gui, kind ModalKind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

func (p *ModalView) SetLines(g1 *gocui.Gui, lines []string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Lines = lines
			return nil
		},
	)
}

func (p *ModalView) SetTitle(g1 *gocui.Gui, title string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Title = title
			return nil
		},
	)
}

func (p *ModalView) SetCursor(g1 *gocui.Gui, pos int) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Cursor = pos
			if len(p.Lines) > 0 {
				p.Selected = p.Lines[p.Cursor]
			}
			p.Selected = ""
			return nil
		},
	)
}

func (p *ModalView) SetSelected(g1 *gocui.Gui, selected string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Selected = selected
			return nil
		},
	)
}
