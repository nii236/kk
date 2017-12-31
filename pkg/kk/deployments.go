package k

import (
	"github.com/jroimartin/gocui"
	appsv1 "k8s.io/api/apps/v1beta1"
)

// DeploymentEntities contains the data for deployments from the API
type DeploymentEntities struct {
	Cursor         int
	Filter         string
	FilterKind     string
	Selected       string
	Deployments    *appsv1.DeploymentList `json:"-"`
	SendingRequest bool
	Size           int
}

// DeploymentFilter is a collection function that filters deployments based on a predicate
func DeploymentFilter(vs []appsv1.Deployment, f func(appsv1.Deployment) bool) []appsv1.Deployment {
	vsf := make([]appsv1.Deployment, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// ClearFilter updates the DeploymentEntities state with an empty filter
func (e *DeploymentEntities) ClearFilter(g1 *gocui.Gui) {
	Debugln("Deployments: ClearFilter")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = ""
			return nil
		},
	)
}

// SetFilter updates the DeploymentEntities state with a new filter
func (e *DeploymentEntities) SetFilter(g1 *gocui.Gui, filter string) {
	Debugln("Deployments: SetFilter")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = filter
			return nil
		},
	)
}

// SetCursor updates the DeploymentEntities state with a new cursor position (absolute)
func (e *DeploymentEntities) SetCursor(g1 *gocui.Gui, pos int) {
	Debugln("Deployments: SetCursor")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Cursor = pos
			return nil
		},
	)
}

// CursorMove updates the DeploymentEntities state with a new cursor position (delta)
func (e *DeploymentEntities) CursorMove(g1 *gocui.Gui, delta int) {
	Debugln("Deployments: CursorMove")
	g1.Update(
		func(g *gocui.Gui) error {
			filteredDeployments := DeploymentFilter(e.Deployments.Items, func(pod appsv1.Deployment) bool {
				if e.Filter == "" {
					return true
				}
				if pod.Namespace == e.Filter {
					return true
				}
				return false
			})
			if len(filteredDeployments) < 2 {
				return nil
			}

			e.Cursor = e.Cursor + delta
			switch {
			case e.Cursor < 1:
				e.Cursor = 1
			case e.Cursor > len(filteredDeployments):
				e.Cursor = len(filteredDeployments)
			}

			e.Selected = filteredDeployments[e.Cursor-1].Name
			return nil
		},
	)
}

// LoadDeploymentData updates the DeploymentEntities state with new data
func (e *DeploymentEntities) LoadDeploymentData(g1 *gocui.Gui, deployments *appsv1.DeploymentList) {
	Debugln("Deployments: LoadDeploymentData")
	g1.Update(
		func(g *gocui.Gui) error {
			e.Size = len(deployments.Items)
			e.Deployments = deployments
			e.SendingRequest = false
			return nil
		},
	)
}
