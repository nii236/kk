package actions

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
)

func LoadManual(client k8s.ClientSet, s *k.State) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		k.Debugln(g, "Load: Manual")
		loader := Load(g, client, s)
		g.Update(loader)
		return nil
	}
}

func Load(g *gocui.Gui, client k8s.ClientSet, s *k.State) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {

		// Load Namespaces
		namespaces, err := client.GetNamespaces()
		if err != nil {
			s.Entities.Errors.Append(g, err)
			return err
		}

		s.Entities.Namespaces.LoadNamespaces(g, namespaces)

		if s.UI.Modal.Kind == k.KindNamespaces {
			nsLines := []string{}
			for _, ns := range namespaces.Items {
				nsLines = append(nsLines, ns.Name)
			}
			s.UI.Modal.SetLines(g, nsLines)
		}

		// Load Pods
		data, err := client.GetPods("")
		if err != nil {
			s.Entities.Errors.Append(g, err)
		}
		s.Entities.Pods.LoadPodData(g, data)
		podList := [][]string{}
		if s.UI.Table.Kind == k.KindPods {
			for _, pod := range data.Items {
				podList = append(podList, k.PodLineHelper(pod))
			}
		}

		// Load Resources

		// Load Deployments
		k.Debugln(g, fmt.Sprintf("Load (Auto): %d Pods, %d Namespaces.", len(data.Items), len(namespaces.Items)))
		return nil
	}
}
