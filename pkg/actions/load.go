package actions

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/common"
	"github.com/nii236/k/pkg/components/state"
	"github.com/nii236/k/pkg/k8s"
)

func LoadMock(client k8s.ClientSet, s *state.Widget) func(g *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, _ *gocui.View) error {
		// Load Pods
		data, err := client.GetPods("kube-system")
		if err != nil {
			s.State.Entities.Errors.Append(g, err)
		}
		s.State.Entities.Pods.LoadPodData(g, data)
		podList := [][]string{}
		if s.State.UI.Table.Kind == k.KindPods {
			for _, pod := range data.Items {
				podList = append(podList, common.PodLineHelper(pod))
			}

			s.State.UI.Table.SetLines(g, podList)
			s.State.UI.Table.SetHeaders(g, k.PodListHeaders)
		}

		// Load Resources

		// Load Namespaces
		namespaces, err := client.GetNamespaces()
		if err != nil {
			s.State.Entities.Errors.Append(g, err)
			return err
		}
		s.State.Entities.Namespaces.LoadNamespaces(g, namespaces)

		if s.State.UI.Modal.Kind == k.KindNamespaces {
			nsLines := []string{}
			for _, ns := range namespaces.Items {
				nsLines = append(nsLines, ns.Name)
			}
			s.State.UI.Modal.SetLines(g, nsLines)
		}
		// Load Deployments

		return nil
	}
}
