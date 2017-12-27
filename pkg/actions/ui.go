package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// Global action: Toggle debug
func ToggleViewDebug(s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenDebug {
			s.UI.SetTableActive(g)
			return nil
		}
		s.UI.SetDebugActive(g)
		return nil

	}
}

func ShowErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		lines := s.Entities.Errors.Lines
		s.UI.SetModalActive(g)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Errors")
		return nil
	}
}

func HideErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.UI.SetTableActive(g)
		return nil
	}
}

func ToggleState(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenState {
			s.UI.SetTableActive(g)
			return nil
		}
		s.UI.SetStateActive(g)
		return nil
	}
}

func ToggleResources(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetTableActive(g)
			return nil
		}
		k.Debugln(g, "Toggle: resources")
		lines := []string{k.KindPods.String(), k.KindNamespaces.String()}
		s.UI.SetModalActive(g)
		s.UI.Modal.SetModalKind(g, k.KindResources)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Resources")
		return nil

	}
}
func ToggleNamespaces(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetTableActive(g)
			return nil
		}
		k.Debugln(g, "Toggle: namespaces")
		lines := []string{}
		for _, ns := range s.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, ns.ObjectMeta.Name)
		}
		s.UI.SetModalActive(g)
		s.UI.Modal.SetModalKind(g, k.KindNamespaces)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Namespaces")
		return nil
	}
}

func PageUp(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Page Up")
		s.UI.CursorMove(g, -10)
		return nil
	}
}

func Prev(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Up")
		s.UI.CursorMove(g, -1)
		return nil
	}
}

func PageDown(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Page Down")
		s.UI.CursorMove(g, 10)
		return nil
	}
}

func Next(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Down")
		s.UI.CursorMove(g, 1)
		return nil
	}
}
func AcknowledgeErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.Entities.Errors.Acknowledge(g)
		s.UI.SetTableActive(g)
		return nil
	}
}
