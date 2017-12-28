package debug

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// Widget represents the debug widget
type Widget struct {
	Name  string
	State *k.State
}

// New returns a new debug widget
func New(name string, initialState *k.State) *Widget {
	return &Widget{
		Name:  name,
		State: initialState,
	}
}

// Layout for the debug widget
func (st *Widget) Layout(g *gocui.Gui) error {
	w, h := g.Size()
	v, err := g.SetView(k.ScreenDebug.String(), 0, 3, w-1, h-4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	if st.State.UI.ActiveScreen != k.ScreenDebug {
		g.SetViewOnBottom(k.ScreenDebug.String())
		return nil
	}

	g.SetViewOnTop(k.ScreenDebug.String())
	v.Title = "Debug"
	v.Autoscroll = true
	return nil
}
