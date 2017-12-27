package k

import (
	"time"

	"github.com/jroimartin/gocui"
)

type EntitiesReducer struct {
	Pods        *PodEntities
	Debug       *DebugEntities
	Errors      *ErrorEntities
	Namespaces  *NamespaceEntities
	Resources   *ResourceEntities
	Deployments *DeploymentEntities
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
