package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/components/state"
	"github.com/nii236/k/pkg/k"
)

func HandleEnter(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenModal {
			switch s.State.UI.Modal.Kind {
			case k.KindResources:
				resource := s.State.UI.Modal.Selected
				s.State.UI.Table.SelectResource(g, resource)
				s.State.UpdateTable(g, k.Kind(resource))
				s.State.UI.SetTableActive(g)
			case k.KindNamespaces:
				selected := s.State.UI.Modal.Selected
				s.State.UI.Table.SetFilter(g, selected)
				s.State.UpdateTable(g, k.Kind(k.KindPods))
				s.State.UI.SetTableActive(g)
			}
		}
		return nil
	}
}
