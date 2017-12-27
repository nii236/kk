package k

import (
	"github.com/jroimartin/gocui"
	"k8s.io/api/core/v1"
)

// PodEntities represents the Pod data
type PodEntities struct {
	Cursor         int
	Filter         string
	FilterKind     string
	Pods           *v1.PodList `json:"-"`
	SendingRequest bool
	Size           int
}

// PodFilter is a collection function that filters pods based on a predicate
func PodFilter(vs []v1.Pod, f func(v1.Pod) bool) []v1.Pod {
	vsf := make([]v1.Pod, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// ClearFilter updates the PodEntities state with an empty filter
func (e *PodEntities) ClearFilter(g1 *gocui.Gui) {
	Debugln("Pods: ClearFilter")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = ""
			return nil
		},
	)
}

// SetFilter updates the PodEntities state with a new filter
func (e *PodEntities) SetFilter(g1 *gocui.Gui, filter string) {
	Debugln("Pods: SetFilter")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = filter
			return nil
		},
	)
}

// SetCursor updates the PodEntities state with a new cursor position
func (e *PodEntities) SetCursor(g1 *gocui.Gui, pos int) {
	Debugln("Pods: SetCursor")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Cursor = pos
			return nil
		},
	)
}

// CursorMove updates the PodEntities state with a new filter
func (e *PodEntities) CursorMove(g1 *gocui.Gui, delta int) {
	Debugln("Pods: CursorMove")
	g1.Update(
		func(g *gocui.Gui) error {
			filteredPods := PodFilter(e.Pods.Items, func(pod v1.Pod) bool {
				if e.Filter == "" {
					return true
				}
				if pod.Namespace == e.Filter {
					return true
				}
				return false
			})
			if len(filteredPods) < 2 {
				return nil
			}

			e.Cursor = e.Cursor + delta
			switch {
			case e.Cursor < 1:
				e.Cursor = 1
			case e.Cursor > len(filteredPods):
				e.Cursor = len(filteredPods)
			}

			return nil
		},
	)
}

// LoadPodData updates the PodEntities state with a new filter
func (e *PodEntities) LoadPodData(g1 *gocui.Gui, pods *v1.PodList) {
	Debugln("Pods: LoadPodData")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Pods = pods
			e.SendingRequest = false
			e.Size = len(pods.Items)
			return nil
		},
	)
}
