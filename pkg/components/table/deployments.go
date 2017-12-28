package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// Lines returns the list of pods for table rendering
func (r *DeploymentRenderer) Lines(s *k.State) [][]string {
	lines := [][]string{}
	for _, deployment := range s.Entities.Deployments.Deployments.Items {
		lines = append(lines, k.DeploymentLineHelper(deployment))
	}

	if s.Entities.Deployments.Filter != "" {
		lines = filter(lines, func(str string) bool {
			if str == s.Entities.Deployments.Filter {
				return true
			}
			return false
		})
	}
	return lines
}

// Cursor returns the cursor position
func (r *DeploymentRenderer) Cursor(s *k.State) int {
	return s.Entities.Deployments.Cursor
}

// Origin returns the origin position
func (r *DeploymentRenderer) Origin(s *k.State, v *gocui.View) (int, int) {
	_, vy := v.Size()
	if s.Entities.Deployments.Cursor < vy {
		return 0, 0
	}
	return 0, s.Entities.Deployments.Cursor - vy + 1
}

// Headers returns the headers
func (r *DeploymentRenderer) Headers(s *k.State) []string {
	return k.PodListHeaders
}

// DeploymentRenderer implements the Renderer interface for use in tables
type DeploymentRenderer struct{}

// NewDeploymentRenderer returns a new DeploymentRenderer to be injected into the new table
func NewDeploymentRenderer() *DeploymentRenderer {
	return &DeploymentRenderer{}
}
