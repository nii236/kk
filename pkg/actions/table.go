package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/k8s"
	"github.com/nii236/kk/pkg/kk"
	"k8s.io/api/core/v1"
)

// TableClearFilter returns a function that will clear the filters for entities displayed in a table
func TableClearFilter(s *k.State) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v2 *gocui.View) error {
		k.Debugln("Table: Clear filter")
		s.Entities.Pods.SetFilter(g, "")
		s.Entities.Deployments.SetFilter(g, "")
		return nil
	}
}

// TableCursorMove returns a function that will move cursors for entities displayed in a table
func TableCursorMove(s *k.State, delta int) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		switch s.UI.Table.Kind {
		case k.KindTableNamespaces:
			s.Entities.Namespaces.CursorMove(g, delta)
		case k.KindTablePods:
			s.Entities.Pods.CursorMove(g, delta)
		case k.KindTableDeployments:
			s.Entities.Deployments.CursorMove(g, delta)
		default:
			k.Errorln("TableCursorMove: Unsupported kind", s.UI.Table.Kind)
		}

		return nil
	}
}

// HandleTableEnter runs when pressing enter while focused on a Table
func HandleTableEnter(s *k.State, c k8s.ClientSet) func(*gocui.Gui, *gocui.View) error {
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
func HandleTableDelete(s *k.State, c k8s.ClientSet) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		k.Debugln("Table: Delete")
		switch s.UI.Table.Kind {
		case k.KindTablePods:
			_, y := v.Cursor()
			line, err := v.Line(y)
			if err != nil {
				k.Errorln(err)
				return err
			}
			podName, err := k.PodNameFromLine(line)
			if err != nil {
				k.Errorln(err)
				return err
			}
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
