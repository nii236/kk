package modal

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

type Widget struct {
	Size SizeEnum
}

type SizeEnum int

const (
	Small SizeEnum = iota + 1
	Medium
	Large
)

func New(name string, size SizeEnum) *Widget {
	return &Widget{
		Size: size,
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
	v, err := g.SetView(k.ScreenModal.String(), x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	store, err := k.JSONToState(g)
	if err != nil {
		fmt.Fprint(v, err)
		return nil
	}
	if store.UI.ActiveScreen != k.ScreenModal {
		g.SetViewOnBottom(k.ScreenModal.String())
		return nil
	}

	g.SetCurrentView(k.ScreenModal.String())
	g.SetViewOnTop(k.ScreenModal.String())
	v.Clear()
	v.Highlight = true
	v.Title = store.UI.Modal.Title
	v.SetCursor(0, store.UI.Modal.Cursor)
	lines := store.UI.Modal.Lines
	v.Write([]byte(strings.Join(lines, "\n")))
	return nil
}
