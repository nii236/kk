package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

func HandleModalEnter(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln(g, "Modal: Pressed enter")
		if s.UI.ActiveScreen == k.ScreenModal {
			switch s.UI.Modal.Kind {
			case k.KindResources:
				resource := s.UI.Modal.Selected
				s.UI.Table.SelectResource(g, resource)
				s.UI.SetTableActive(g)
			case k.KindNamespaces:
				selected := s.UI.Modal.Selected
				s.Entities.Pods.SetFilter(g, selected)
				s.UI.SetTableActive(g)
			}
		}
		return nil
	}
}

func HandleTableEnter(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln(g, "Table: pressed enter")
		return nil
	}
}
