package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/kk/pkg/kk"
)

// Lines returns the list of namespaces for table rendering
func (r *NamespaceRenderer) Lines(s *k.State) [][]string {
	lines := [][]string{}
	for _, ns := range s.Entities.Namespaces.Namespaces.Items {
		lines = append(lines, k.NamespaceLineHelper(ns))
	}
	return lines
}

// Cursor returns the cursor position
func (r *NamespaceRenderer) Cursor(s *k.State) int {
	return s.Entities.Namespaces.Cursor
}

// Origin returns the origin position
func (r *NamespaceRenderer) Origin(s *k.State, v *gocui.View) (int, int) {
	_, vy := v.Size()
	if s.Entities.Namespaces.Cursor < vy {
		return 0, 0
	}
	return 0, s.Entities.Namespaces.Cursor - vy + 1
}

// Headers returns the headers
func (r *NamespaceRenderer) Headers(s *k.State) []string {
	return k.NamespaceListHeaders
}

// NamespaceRenderer implements the Renderer interface for use in tables
type NamespaceRenderer struct{}

// NewNamespaceRenderer returns a new NamespaceRenderer to be injected into the new table
func NewNamespaceRenderer() *NamespaceRenderer {
	return &NamespaceRenderer{}
}
