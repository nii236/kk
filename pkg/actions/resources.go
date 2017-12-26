package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/components/state"
)

func SelectResource(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		resource := s.State.UI.Modal.Selected
		switch resource {
		case k.KindNamespaces.String():
			s.State.Entities.Debug.Append(g, "Show NS")
			switchToNamespace := ShowNamespaceList(s)
			switchToNamespace(g, v2)
		case k.KindPods.String():
			s.State.Entities.Debug.Append(g, "Show NS")
			switchToPods := ShowPodList(s)
			switchToPods(g, v2)
		}
		return nil
	}
}
