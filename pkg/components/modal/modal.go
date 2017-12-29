package modal

import (
	"github.com/nii236/k"

	"github.com/jroimartin/gocui"
)

type Widget struct {
	Name     string
	Size     SizeEnum
	Selected int
	Visible  bool
	State    *k.State
}

type SizeEnum int

const (
	Small SizeEnum = iota + 1
	Medium
	Large
)

func New(name string, size SizeEnum, state *k.State) *Widget {
	return &Widget{
		Name:     name,
		Size:     size,
		Selected: 0,
		Visible:  false,
		State:    state,
	}
}

// Layout for the tablewidget
func (st *Widget) Layout(g *gocui.Gui) error {

	w, h := g.Size()
	modalWidth := 0
	modalHeight := 0
	switch st.Size {
	case Small:
		modalWidth = w * 1 / 8
		modalHeight = h * 1 / 8
	case Medium:
		modalWidth = w * 1 / 4
		modalHeight = h * 1 / 4
	case Large:
		modalWidth = w * 1 / 2
		modalHeight = h * 1 / 2
	default:
		modalWidth = w
		modalHeight = h
	}

	x0 := w/2 - modalWidth/2
	x1 := w/2 + modalWidth/2
	y0 := h/2 - modalHeight/2
	y1 := h/2 + modalHeight/2
	v, err := g.SetView(st.Name, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Highlight = true
	if !st.Visible {
		g.SetViewOnBottom(st.Name)
		return nil
	}
	g.SetViewOnTop(st.Name)
	v.Title = st.Name
	return nil
}

func Toggle(w *Widget) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(w.Name)
		w.Visible = !w.Visible
		return nil
	}
}
