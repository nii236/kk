package ui

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/common"
	"github.com/nii236/k/pkg/components/state"
	"github.com/nii236/k/pkg/k8s"
)

var DEBUG_DISPLAYED bool = false

// Global action: Toggle debug
func ActionToggleViewDebug(s *state.Widget) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenDebug {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.UI.SetDebugActive(g)
		return nil

	}
}

func ActionLoadMock(client k8s.ClientSet, s *state.Widget) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		data, err := client.GetPods("kube-system")
		if err != nil {
			displayError(g, err)
		}
		s.State.Entities.Pods.LoadPodData(g, data)
		podList := [][]string{}
		for _, pod := range data.Items {
			name := pod.ObjectMeta.Name
			restarts := common.ColumnHelperRestarts(pod)
			age := common.ColumnHelperAge(pod)
			ready := common.ColumnHelperReady(pod)
			status := common.ColumnHelperStatus(pod)
			podList = append(podList, []string{name, restarts, age, ready, status})
		}

		s.State.UI.Table.SetTableLines(g, podList)

		return nil
	}

}

func ActionShowPods(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		s.State.UI.SetTableActive(g)
		s.State.UI.Table.SetTableKind(g, k.KindPods)
		return nil
	}
}

func ActionToggleState(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenState {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.UI.SetStateActive(g)
		return nil
	}
}

func ActionToggleResources(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenModal {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.Entities.Debug.Append(g, "Toggle screen to: "+"resources")
		lines := []string{"HI", "BYE"}
		s.State.UI.SetModalActive(g)
		s.State.UI.Modal.SetModalLines(g, lines)
		s.State.UI.Modal.SetModalTitle(g, "Resources")
		return nil

	}
}
func ActionToggleNamespaces(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		if s.State.UI.ActiveScreen == k.ScreenModal {
			s.State.UI.SetTableActive(g)
			return nil
		}
		s.State.Entities.Debug.Append(g, "Toggle screen to: "+"namespaces")
		lines := []string{"HI2", "BYE2"}
		s.State.UI.SetModalActive(g)
		s.State.UI.Modal.SetModalLines(g, lines)
		s.State.UI.Modal.SetModalTitle(g, "Namespaces")
		return nil
	}
}

func ActionPrev(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Move Up")
		switch s.State.UI.ActiveScreen {
		case k.ScreenDebug:
			s.State.UI.Debug.CursorUp(g)
		case k.ScreenTable:
			s.State.UI.Table.CursorUp(g)
		case k.ScreenModal:
			s.State.UI.Modal.CursorUp(g)
		case k.ScreenState:
			s.State.UI.State.CursorUp(g)
		}
		return nil
	}
}

func ActionNext(s *state.Widget) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		s.State.Entities.Debug.Append(g, "Move Down")
		switch s.State.UI.ActiveScreen {
		case k.ScreenDebug:
			s.State.UI.Debug.CursorDown(g)
		case k.ScreenTable:
			s.State.UI.Table.CursorDown(g)
		case k.ScreenModal:
			s.State.UI.Modal.CursorDown(g)
		case k.ScreenState:
			s.State.UI.State.CursorDown(g)
		}
		return nil
	}
}

func HideError(g *gocui.Gui, _ *gocui.View) error {
	hideErrorPopup(g)
	return nil
}
