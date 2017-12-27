package k

import (
	"github.com/jroimartin/gocui"
	"k8s.io/api/core/v1"
)

type NamespaceEntities struct {
	Cursor         int
	Loaded         bool
	Filter         string
	FilterKind     string
	Selected       string
	Namespaces     *v1.NamespaceList
	SendingRequest bool
}

func (pr *NamespaceEntities) LoadNamespaces(g1 *gocui.Gui, ns *v1.NamespaceList) {
	g1.Update(
		func(g *gocui.Gui) error {
			pr.Loaded = true
			pr.Namespaces = ns
			pr.SendingRequest = false
			return nil
		},
	)
}

func (e *NamespaceEntities) CursorMove(g *gocui.Gui, delta int) {
	if len(e.Namespaces.Items) < 2 {
		return
	}
	e.Cursor = e.Cursor + delta
	switch {
	case e.Cursor < 1:
		e.Cursor = 1
	case e.Cursor > len(e.Namespaces.Items):
		e.Cursor = len(e.Namespaces.Items)
	}

	e.Selected = e.Namespaces.Items[e.Cursor-1].Name
}
