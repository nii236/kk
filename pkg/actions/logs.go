package actions

import (
	"bytes"
	"fmt"
	"strconv"
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
			k.Debugln("logs: Get containers for pod " + s.Entities.Pods.Selected)
			podName := s.Entities.Pods.Selected
			podToFetch := &v1.Pod{}
			for _, pod := range s.Entities.Pods.Pods.Items {
				if podName == pod.Name {
					podToFetch = &pod
					break
				}
			}

			s.UI.SetActiveScreen(g, k.ScreenModal)

			containers := c.GetPodContainers(podToFetch.Name, podToFetch.Namespace)
			s.UI.Modal.SetCursor(g, 0)
			s.UI.Modal.SetSelected(g, "")
			k.Debugln("Pod: " + podToFetch.Name)
			k.Debugln("Namespace: " + podToFetch.Namespace)
			k.Debugln("Containers: " + strings.Join(containers, ", "))
			k.Debugln("len(containers)" + strconv.Itoa(len(containers)))
			if len(containers) > 0 {
				s.UI.Modal.SetSelected(g, containers[0])
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

		selectedPod := s.Entities.Pods.Selected
		podToFetch := &v1.Pod{}
		for _, pod := range s.Entities.Pods.Pods.Items {
			if selectedPod == pod.Name {
				podToFetch = &pod
				break
			}
		}
		selectedContainer := s.UI.Modal.Selected
		containerLabel := selectedContainer
		if selectedContainer == "" {
			containerLabel = "default"
		}
		k.Debugln("selectedPod: " + selectedPod)
		k.Debugln("selectedContainer: " + selectedContainer)
		k.Debugln("namespace: " + podToFetch.Namespace)

		var buf bytes.Buffer
		c.GetPodContainerLogs(podToFetch.Name, selectedContainer, podToFetch.Namespace, &buf)
		s.UI.Modal.SetKind(g, k.KindModalContainerLogs)
		s.UI.Modal.SetTitle(g, fmt.Sprintf("%s: %s -> %s -> %s", k.KindModalContainerLogs, podToFetch.Namespace, podToFetch.Name, containerLabel))
		s.UI.Modal.SetLines(g, strings.Split(buf.String(), "\n"))
		s.UI.Modal.SetSize(g, k.ModalSizeExtraLarge)
		return nil
	}
}
