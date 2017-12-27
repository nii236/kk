package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
)

func LoadAuto(client k8s.ClientSet, s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		// k.Debugln("Load: Auto")
		loader := load(g, client, s)
		g.Update(loader)
		return nil
	}
}

func LoadManual(client k8s.ClientSet, s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln("Load: Manual")
		loader := load(g, client, s)
		g.Update(loader)
		return nil
	}
}

func load(g *gocui.Gui, client k8s.ClientSet, s *k.State) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {

		// Load Namespaces
		namespaces, err := client.GetNamespaces()
		if err != nil {
			k.Errorln(err)
			return err
		}

		s.Entities.Namespaces.LoadNamespaces(g, namespaces)

		if s.UI.Modal.Kind == k.KindModalNamespaces {
			nsLines := []string{}
			for _, ns := range namespaces.Items {
				nsLines = append(nsLines, ns.Name)
			}
			s.UI.Modal.SetLines(g, nsLines)
		}

		// Load Pods
		pods, err := client.GetPods("")
		if err != nil {
			k.Errorln(err)
			return err
		}
		s.Entities.Pods.LoadPodData(g, pods)
		if s.UI.Table.Kind == k.KindTablePods {
			if s.Entities.Pods.Cursor == 0 {
				s.Entities.Pods.SetCursor(g, 1)
				s.Entities.Pods.SetSelected(g, "")
				if len(pods.Items) > 0 {
					s.Entities.Pods.SetSelected(g, pods.Items[0].Name)
				}
			}
		}

		// Load Deployments
		deployments, err := client.GetDeployments("")
		if err != nil {
			k.Errorln(err)
			return err
		}
		s.Entities.Deployments.LoadDeploymentData(g, deployments)
		if s.UI.Table.Kind == k.KindTableDeployments {
			if s.Entities.Deployments.Cursor == 0 {
				s.Entities.Deployments.SetCursor(g, 1)
				s.Entities.Deployments.SetSelected(g, "")
				if len(deployments.Items) > 0 {
					s.Entities.Deployments.SetSelected(g, deployments.Items[0].Name)
				}
			}
		}

		// Load Resources

		// k.Debugln(fmt.Sprintf("Load: %d Pods, %d Namespaces.", len(data.Items), len(namespaces.Items)))
		return nil
	}
}
