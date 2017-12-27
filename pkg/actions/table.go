package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// TableClearFilter returns a function that will clear the filters for entities displayed in a table
func TableClearFilter(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Table: Clear filter")
		s.Entities.Pods.SetFilter(g, "")
		s.Entities.Deployments.SetFilter(g, "")
		return nil
	}
}

// TableCursorMove returns a function that will move cursors for entities displayed in a table
func TableCursorMove(s *k.State, delta int) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		switch s.UI.Table.Kind {
		case k.KindTableNamespaces:
			s.Entities.Namespaces.CursorMove(g, delta)
		case k.KindTablePods:
			s.Entities.Pods.CursorMove(g, delta)
		case k.KindTableDeployments:
			s.Entities.Deployments.CursorMove(g, delta)
		default:
			k.Errorln("TableCursorMove: Unsupported kind", s.UI.Table.Kind)
		}

		return nil
	}
}
