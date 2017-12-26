package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/common"
	"github.com/nii236/k/pkg/components/state"
)

// Global action: Toggle debug
func ToggleViewDebug(s *state.Widget) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenDebug {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.UI.SetDebugActive(g)
		return nil

	}
}

func ShowErrors(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		lines := s.State.Entities.Errors.Lines
		s.State.UI.SetModalActive(g)
		s.State.UI.Modal.SetLines(g, lines)
		s.State.UI.Modal.SetTitle(g, "Errors")
		return nil
	}
}

func HideErrors(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.UI.SetTableActive(g)
		return nil
	}
}

func ShowPodList(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.UI.SetTableActive(g)
		s.State.UI.Table.SetKind(g, k.KindPods)
		podList := [][]string{}
		for _, pod := range s.State.Entities.Pods.Pods.Items {
			podList = append(podList, common.PodLineHelper(pod))
		}
		s.State.UI.Table.SetLines(g, podList)
		s.State.UI.Table.SetHeaders(g, k.PodListHeaders)
		return nil
	}
}

func ShowNamespaceList(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.UI.SetTableActive(g)
		s.State.UI.Table.SetKind(g, k.KindNamespaces)
		lines := [][]string{}
		for _, ns := range s.State.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, []string{ns.Name})
		}
		s.State.UI.Table.SetLines(g, lines)
		s.State.UI.Table.SetHeaders(g, k.NamespaceListHeaders)
		return nil
	}
}

func ToggleState(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenState {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.UI.SetStateActive(g)
		return nil
	}
}

func ToggleResources(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenModal {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.Entities.Debug.Append(g, "Toggle screen to: "+"resources")
		lines := []string{k.KindPods.String(), k.KindNamespaces.String()}
		s.State.UI.SetModalActive(g)
		s.State.UI.Modal.SetModalKind(g, k.KindResources)
		s.State.UI.Modal.SetLines(g, lines)
		s.State.UI.Modal.SetTitle(g, "Resources")
		return nil

	}
}
func ToggleNamespaces(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenModal {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.Entities.Debug.Append(g, "Toggle screen to: "+"namespaces")
		lines := []string{}
		for _, ns := range s.State.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, ns.ObjectMeta.Name)
		}
		s.State.UI.SetModalActive(g)
		s.State.UI.Modal.SetModalKind(g, k.KindNamespaces)
		s.State.UI.Modal.SetLines(g, lines)
		s.State.UI.Modal.SetTitle(g, "Namespaces")
		return nil
	}
}

func PageUp(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Page Up")
		s.State.UI.CursorMove(g, -10)
		return nil
	}
}

func Prev(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Up")
		s.State.UI.CursorMove(g, -1)
		return nil
	}
}

func PageDown(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Page Down")
		s.State.UI.CursorMove(g, 10)
		return nil
	}
}

func Next(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Down")
		s.State.UI.CursorMove(g, 1)
		return nil
	}
}
func AcknowledgeErrors(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Errors.Acknowledge(g)
		s.State.UI.SetTableActive(g)
		return nil
	}
}
