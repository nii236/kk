package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/components/state"
	"github.com/nii236/k/pkg/k"
)

func TableClearFilter(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.Entities.Pods.SetFilter(g, "")
		if s.State.UI.Table.Kind == k.KindPods {
		}
		return nil
	}
}

func TableCursorMoveUp(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Up")
		if s.State.UI.Table.Kind == k.KindNamespaces {
			s.State.Entities.Namespaces.CursorMove(g, -1)
		}
		if s.State.UI.Table.Kind == k.KindPods {
			k.Debugln("Table move up")
			s.State.Entities.Pods.CursorMove(g, -1)
		}
		return nil
	}
}

func TableCursorMoveDown(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Down")
		if s.State.UI.Table.Kind == k.KindNamespaces {
			s.State.Entities.Namespaces.CursorMove(g, 1)
		}
		if s.State.UI.Table.Kind == k.KindPods {
			k.Debugln("Table move down")
			s.State.Entities.Pods.CursorMove(g, 1)
		}
		return nil
	}
}
