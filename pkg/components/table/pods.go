package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/kk"
)

// Lines returns the list of pods for table rendering
func (r *PodRenderer) Lines(s *k.State) [][]string {
	lines := [][]string{}
	for _, pod := range s.Entities.Pods.Pods.Items {
		lines = append(lines, k.PodLineHelper(pod))
	}

	if s.Entities.Pods.Filter != "" {
		lines = filter(lines, func(str string) bool {
			if str == s.Entities.Pods.Filter {
				return true
			}
			return false
		})
	}
	return lines
}

// Cursor returns the cursor position
func (r *PodRenderer) Cursor(s *k.State) int {
	return s.Entities.Pods.Cursor
}

// Origin returns the origin position
func (r *PodRenderer) Origin(s *k.State, v *gocui.View) (int, int) {
	_, vy := v.Size()
	if s.Entities.Pods.Cursor < vy {
		return 0, 0
	}
	return 0, s.Entities.Pods.Cursor - vy + 1
}

// Headers returns the headers
func (r *PodRenderer) Headers(s *k.State) []string {
	return k.PodListHeaders
}

// PodRenderer implements the Renderer interface for use in tables
type PodRenderer struct{}

// NewPodRenderer returns a new PodRenderer to be injected into the new table
func NewPodRenderer() *PodRenderer {
	return &PodRenderer{}
}
