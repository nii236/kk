package k

import (
	"time"

	"github.com/jroimartin/gocui"
	v1 "k8s.io/api/core/v1"
)

type EntitiesReducer struct {
	Pods  *PodEntities
	Debug *DebugEntities
}

type DebugEntities struct {
	Lines []interface{}
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

type PodEntities struct {
	Loaded         bool
	Pods           *v1.PodList
	SendingRequest bool
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
