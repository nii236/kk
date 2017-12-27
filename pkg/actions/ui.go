package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// Global action: Toggle debug
func ToggleViewDebug(s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenDebug {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		s.UI.SetActiveScreen(g, k.ScreenDebug)
		return nil

	}
}

func ShowErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		lines := s.Entities.Errors.Lines
		s.UI.SetActiveScreen(g, k.ScreenModal)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Errors")
		return nil
	}
}

func HideErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.UI.SetActiveScreen(g, k.ScreenTable)
		return nil
	}
}

func ToggleResources(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		k.Debugln("Toggle: resources")
		lines := s.Entities.Resources.Resources
		s.UI.SetActiveScreen(g, k.ScreenModal)
		s.UI.Modal.SetCursor(g, 0)
		s.UI.Modal.SetKind(g, k.KindModalResources)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Resources")
		s.UI.Modal.SetSize(g, k.ModalSizeMedium)
		return nil

	}
}
func ToggleNamespaces(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.UI.ActiveScreen == k.ScreenModal {
			s.UI.SetActiveScreen(g, k.ScreenTable)
			return nil
		}
		k.Debugln("Toggle: namespaces")
		lines := []string{}
		for _, ns := range s.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, ns.ObjectMeta.Name)
		}
		s.UI.SetActiveScreen(g, k.ScreenModal)
		s.UI.Modal.SetCursor(g, 0)
		s.UI.Modal.SetKind(g, k.KindModalNamespaces)
		s.UI.Modal.SetLines(g, lines)
		s.UI.Modal.SetTitle(g, "Namespaces")
		s.UI.Modal.SetSize(g, k.ModalSizeMedium)
		return nil
	}
}

func PageUp(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.UI.CursorMove(g, -10)
		return nil
	}
}

func Prev(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.UI.CursorMove(g, -1)
		return nil
	}
}

func PageDown(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.UI.CursorMove(g, 10)
		return nil
	}
}

func Next(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.UI.CursorMove(g, 1)
		return nil
	}
}
func AcknowledgeErrors(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.Entities.Errors.Acknowledge(g)
		s.UI.SetActiveScreen(g, k.ScreenTable)
		return nil
	}
}
