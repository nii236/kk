package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/components/state"
)

func ClearFilter(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.UI.Table.ClearFilter(g)
		s.State.UpdateTable(g, k.KindPods)
		return nil
	}
}
