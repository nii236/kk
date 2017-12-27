package k

import (
	"strconv"

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
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = ""
			return nil
		},
	)
}

// SetFilter updates the DeploymentEntities state with a new filter
func (e *DeploymentEntities) SetFilter(g1 *gocui.Gui, filter string) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Filter = filter
			return nil
		},
	)
}

// SetCursor updates the DeploymentEntities state with a new cursor position (absolute)
func (e *DeploymentEntities) SetCursor(g1 *gocui.Gui, pos int) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Cursor = pos
			return nil
		},
	)
}

// SetSelected updates the DeploymentEntities state with a new selection
func (e *DeploymentEntities) SetSelected(g1 *gocui.Gui, selected string) {
	g1.Update(
		func(g *gocui.Gui) error {
			e.Selected = selected
			return nil
		},
	)
}

// CursorMove updates the DeploymentEntities state with a new cursor position (delta)
func (e *DeploymentEntities) CursorMove(g1 *gocui.Gui, delta int) {
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

			Debugln("Filtered Deployments: " + strconv.Itoa(len(filteredDeployments)))
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
func (pr *DeploymentEntities) LoadDeploymentData(g1 *gocui.Gui, deployments *appsv1.DeploymentList) {
	g1.Update(
		func(g *gocui.Gui) error {
			pr.Size = len(deployments.Items)
			pr.Deployments = deployments
			pr.SendingRequest = false
			return nil
		},
	)
}
