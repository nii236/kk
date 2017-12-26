package k

import (
	"strconv"
	"time"

	"github.com/jroimartin/gocui"
	v1 "k8s.io/api/core/v1"
)

type EntitiesReducer struct {
	Pods       *PodEntities
	Debug      *DebugEntities
	Errors     *ErrorEntities
	Namespaces *NamespaceEntities
	Resources  *ResourceEntities
}

type ResourceEntities struct {
	Resources []string
}

type ErrorEntities struct {
	Lines        []string
	Acknowledged bool
}
type DebugEntities struct {
	Lines []interface{}
}

type NamespaceEntities struct {
	Cursor         int
	Loaded         bool
	Filter         string
	FilterKind     string
	Selected       string
	Namespaces     *v1.NamespaceList
	SendingRequest bool
}

type PodEntities struct {
	Cursor         int
	Loaded         bool
	Filter         string
	FilterKind     string
	Selected       string
	Pods           *v1.PodList
	SendingRequest bool
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

func (pr *NamespaceEntities) LoadNamespaces(g1 *gocui.Gui, ns *v1.NamespaceList) {
	g1.Update(
		func(g *gocui.Gui) error {
			pr.Loaded = true
			pr.Namespaces = ns
			pr.SendingRequest = false
			return nil
		},
	)
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

			Debugln("Filtered Pods: " + strconv.Itoa(len(filteredPods)))
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

func (e *NamespaceEntities) CursorMove(g *gocui.Gui, delta int) {
	if len(e.Namespaces.Items) < 2 {
		return
	}
	e.Cursor = e.Cursor + delta
	switch {
	case e.Cursor < 1:
		e.Cursor = 1
	case e.Cursor > len(e.Namespaces.Items):
		e.Cursor = len(e.Namespaces.Items)
	}

	e.Selected = e.Namespaces.Items[e.Cursor-1].Name
}

func (d *ErrorEntities) Acknowledge(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			d.Lines = []string{}
			d.Acknowledged = true
			return nil
		},
	)
}

func (d *ErrorEntities) Append(g1 *gocui.Gui, err error) {
	g1.Update(
		func(g *gocui.Gui) error {
			t := time.Now()
			tf := t.Format("2006-01-02 15:04:05")
			val := tf + " > " + err.Error()
			d.Lines = append(d.Lines, val)
			d.Acknowledged = false
			return nil
		},
	)
}

func (d *DebugEntities) Append(g1 *gocui.Gui, val interface{}) {
	g1.Update(
		func(g *gocui.Gui) error {
			t := time.Now()
			tf := t.Format("2006-01-02 15:04:05")
			val = tf + " > " + val.(string)
			d.Lines = append(d.Lines, val)
			return nil
		},
	)

}

func (pr *PodEntities) LoadPodData(g1 *gocui.Gui, pods *v1.PodList) {
	g1.Update(
		func(g *gocui.Gui) error {
			pr.Loaded = true
			pr.Pods = pods
			pr.SendingRequest = false
			return nil
		},
	)
}
