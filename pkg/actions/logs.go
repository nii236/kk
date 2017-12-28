package actions

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
	"k8s.io/api/core/v1"
)

// FetchContainers is a function factory that returns a function that will fetch containers for a pod
func FetchContainers(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if s.UI.Table.Kind == k.KindTablePods {
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
			var podToFetch *v1.Pod
			for _, pod := range s.Entities.Pods.Pods.Items {
				if podName == pod.Name {
					podToFetch = &pod
					break
				}
			}
			if podToFetch == nil {
				err := errors.New("No matching containers found for pod " + podName)
				k.Errorln(err)
				return err
			}

			s.UI.SetActiveScreen(g, k.ScreenModal)

			containers, err := c.GetPodContainers(podToFetch.Name, podToFetch.Namespace)
			if err != nil {
				k.Errorln(err)
			}

			s.UI.Modal.SetKind(g, k.KindModalSelectContainer)
			s.UI.Modal.SetTitle(g, k.KindModalSelectContainer.String())
			s.UI.Modal.SetLines(g, containers)
			s.UI.Modal.SetSize(g, k.ModalSizeSmall)
		}
		return nil
	}
}

// FetchLogs is a function factory that will return a function that fetches logs for a container
func FetchLogs(s *k.State, c k8s.ClientSet) func(g1 *gocui.Gui, _ *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		k.Debugln("logs: Fetch logs for container")
		if s.UI.ActiveScreen != k.ScreenModal &&
			s.UI.Modal.Kind != k.KindModalSelectContainer {
			return nil
		}

		podCursor := s.Entities.Pods.Cursor
		podToFetch := s.Entities.Pods.Pods.Items[podCursor-1]
		k.Debugln("Fetching containers for pod:", podToFetch.Name)

		_, vy := v.Cursor()
		selected, err := v.Line(vy)
		if err != nil {
			k.Errorln(err)
			return err
		}

		containerLabel := selected
		if selected == "" {
			containerLabel = "default"
		}

		var buf bytes.Buffer
		c.GetPodContainerLogs(podToFetch.Name, selected, podToFetch.Namespace, &buf)
		s.UI.Modal.SetKind(g, k.KindModalContainerLogs)
		s.UI.Modal.SetTitle(g, fmt.Sprintf("%s: %s -> %s -> %s", k.KindModalContainerLogs, podToFetch.Namespace, podToFetch.Name, containerLabel))
		s.UI.Modal.SetLines(g, strings.Split(buf.String(), "\n"))
		s.UI.Modal.SetSize(g, k.ModalSizeExtraLarge)
		return nil
	}
}
