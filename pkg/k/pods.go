package k

import (
	"github.com/jroimartin/gocui"
	"k8s.io/api/core/v1"
)

type PodEntities struct {
	Cursor         int
	Filter         string
	FilterKind     string
	Selected       string
	Pods           *v1.PodList `json:"-"`
	SendingRequest bool
	Size           int
}

func PodFilter(vs []v1.Pod, f func(v1.Pod) bool) []v1.Pod {
	vsf := make([]v1.Pod, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func (e *PodEntities) ClearFilter(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = ""
			return nil
		},
	)
}

func (e *PodEntities) SetFilter(g1 *gocui.Gui, filter string) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = filter
			return nil
		},
	)
}

func (e *PodEntities) SetCursor(g1 *gocui.Gui, pos int) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Cursor = pos
			return nil
		},
	)
}

func (p *PodEntities) SetSelected(g1 *gocui.Gui, selected string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Selected = selected
			return nil
		},
	)
}

func (e *PodEntities) CursorMove(g1 *gocui.Gui, delta int) {
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

			e.Selected = filteredPods[e.Cursor-1].Name
			return nil
		},
	)
}

func (pr *PodEntities) LoadPodData(g1 *gocui.Gui, pods *v1.PodList) {
	g1.Update(
		func(g *gocui.Gui) error {
			pr.Pods = pods
			pr.SendingRequest = false
			pr.Size = len(pods.Items)
			return nil
		},
	)
}
