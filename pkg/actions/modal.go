package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/k8s"
	"github.com/nii236/kk/pkg/kk"
)

// ModalCursorMove returns a function that will move cursors for entities displayed in a modal
func ModalCursorMove(s *k.State, delta int) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		s.UI.Modal.CursorMove(g, delta)
		return nil
	}
}

// HandleModalEnter runs when pressing enter while focused on a Modal
func HandleModalEnter(s *k.State, c k8s.ClientSet) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		k.Debugln("Modal: Pressed enter")
		if s.UI.ActiveScreen == k.ScreenModal {
			switch s.UI.Modal.Kind {
			case k.KindModalResources:
				_, y := v.Cursor()
				selected, err := v.Line(y)
				if err != nil {
					k.Errorln(err)
					return err
				}
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.TableKind(selected))
			case k.KindModalNamespaces:
				_, y := v.Cursor()
				selected, err := v.Line(y)
				if err != nil {
					k.Errorln(err)
					return err
				}
				s.Entities.Pods.SetFilter(g, selected)
				s.Entities.Deployments.SetFilter(g, selected)
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.KindTablePods)
			case k.KindModalSelectContainer:
				logFetcher := FetchLogs(s, c)
				logFetcher(g, v)
				s.UI.SetActiveScreen(g, k.ScreenModal)
			case k.KindModalContainerLogs:
				s.UI.SetActiveScreen(g, k.ScreenTable)
			default:
				k.Errorln("Unsupported Modal Kind: " + s.UI.Modal.Kind)
			}
		}
		return nil
	}
}

// HandleModalEsc runs when pressing esc while focused on a Modal
func HandleModalEsc(s *k.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Modal: Pressed esc")
		if s.UI.ActiveScreen == k.ScreenModal {
			switch s.UI.Modal.Kind {
			case k.KindModalResources:
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.KindTablePods)
			case k.KindModalNamespaces:
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.KindTablePods)
			case k.KindModalSelectContainer:
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.KindTablePods)
			case k.KindModalContainerLogs:
				s.UI.SetActiveScreen(g, k.ScreenTable)
				s.UI.Table.SetKind(g, k.KindTablePods)
			default:
				k.Errorln("Unsupported Modal Kind: " + s.UI.Modal.Kind)
			}
		}
		return nil
	}
}
