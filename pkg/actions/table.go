package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
	"k8s.io/api/core/v1"
)

func TableClearFilter(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln(g, "Table: Clear filter")
		s.Entities.Pods.SetFilter(g, "")
		if s.UI.Table.Kind == k.KindPods {
		}
		return nil
	}
}

func TableCursorMoveUp(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Table: Move cursor up")
		if s.UI.Table.Kind == k.KindNamespaces {
			s.Entities.Namespaces.CursorMove(g, -1)
		}
		if s.UI.Table.Kind == k.KindPods {
			s.Entities.Pods.CursorMove(g, -1)
		}
		return nil
	}
}

func TableCursorMoveDown(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		k.Debugln(g, "Table: Move cursor down")
		if s.UI.Table.Kind == k.KindNamespaces {
			s.Entities.Namespaces.CursorMove(g, 1)
		}
		if s.UI.Table.Kind == k.KindPods {
			s.Entities.Pods.CursorMove(g, 1)
		}
		return nil
	}
}

func TableDelete(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		k.Debugln(g, "Table: Delete "+s.Entities.Pods.Selected)
		podName := s.Entities.Pods.Selected
		podToDelete := &v1.Pod{}
		for _, pod := range s.Entities.Pods.Pods.Items {
			if podName == pod.Name {
				podToDelete = &pod
				break
			}
		}
		c.DeletePod(podToDelete.Name, podToDelete.Namespace)
		return nil
	}
}
