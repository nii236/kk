package k

import (
	"github.com/jroimartin/gocui"
	"k8s.io/api/core/v1"
)

// NamespaceEntities contains the data for namespaces
type NamespaceEntities struct {
	Size           int
	Cursor         int
	Filter         string
	FilterKind     string
	Selected       string
	Namespaces     *v1.NamespaceList `json:"-"`
	SendingRequest bool
}

// LoadNamespaces updates the namespace state with a new dataset
func (e *NamespaceEntities) LoadNamespaces(g1 *gocui.Gui, ns *v1.NamespaceList) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Size = len(ns.Items)
			e.Namespaces = ns
			e.SendingRequest = false
			return nil
		},
	)
}

// CursorMove updates the namespace state with a new cursor position (delta)
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
