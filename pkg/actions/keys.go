package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
	"k8s.io/api/core/v1"
)

// HandleModalEnter runs when pressing enter while focused on a Modal
func HandleModalEnter(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Modal: Pressed enter")
		if s.UI.ActiveScreen == k.ScreenModal {
			switch s.UI.Modal.Kind {
			case k.KindModalResources:
				resource := s.UI.Modal.Selected
				s.UI.Table.SelectResource(g, resource)
				s.UI.SetActiveScreen(g, k.ScreenTable)
			case k.KindModalNamespaces:
				selected := s.UI.Modal.Selected
				s.Entities.Pods.SetFilter(g, selected)
				s.Entities.Deployments.SetFilter(g, selected)
				s.UI.SetActiveScreen(g, k.ScreenTable)
			case k.KindModalSelectContainer:
				logFetcher := FetchLogs(s, c)
				logFetcher(g, v2)
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

// HandleDebugEsc runs when pressing esc while focused on a Modal
func HandleDebugEsc(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Modal: Pressed esc")
		s.UI.SetActiveScreen(g, k.ScreenTable)

		return nil
	}
}

// HandleModalEsc runs when pressing esc while focused on a Modal
func HandleModalEsc(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Modal: Pressed esc")
		if s.UI.ActiveScreen == k.ScreenModal {
			switch s.UI.Modal.Kind {
			case k.KindModalResources:
				s.UI.SetActiveScreen(g, k.ScreenTable)
			case k.KindModalNamespaces:
				s.UI.SetActiveScreen(g, k.ScreenTable)
			case k.KindModalSelectContainer:
				s.UI.SetActiveScreen(g, k.ScreenTable)
			case k.KindModalContainerLogs:
				s.UI.SetActiveScreen(g, k.ScreenTable)
			default:
				k.Errorln("Unsupported Modal Kind: " + s.UI.Modal.Kind)
			}
		}
		return nil
	}
}

// HandleTableEnter runs when pressing enter while focused on a Table
func HandleTableEnter(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Table: pressed enter")
		if s.UI.Table.Kind == k.KindTablePods {
			containerFetcher := FetchContainers(s, c)
			containerFetcher(g, v2)
		}
		return nil
	}
}

// HandleTableDelete runs when pressing delete while focused on a Table
func HandleTableDelete(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Table: Delete " + s.Entities.Pods.Selected)
		switch s.UI.Table.Kind {
		case k.KindTablePods:
			podName := s.Entities.Pods.Selected
			podToDelete := &v1.Pod{}
			for _, pod := range s.Entities.Pods.Pods.Items {
				if podName == pod.Name {
					podToDelete = &pod
					break
				}
			}
			c.DeletePod(podToDelete.Name, podToDelete.Namespace)
		default:
			k.Errorln("Table Delete: Unsupported Kind", s.UI.Table.Kind)
		}
		return nil
	}
}
