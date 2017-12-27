package debug

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

type Widget struct {
	Name  string
	State *k.State
}

//ss
func New(name string, initialState *k.State) *Widget {
	return &Widget{
		Name:  name,
		State: initialState,
	}
}

// Layout for the tablewidget
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
	g.SetCurrentView(k.ScreenDebug.String())
	v.Title = "Debug"
	v.Autoscroll = true
	lines := []string{}
	for _, el := range st.State.Entities.Debug.Lines {
		lines = append(lines, el.(string))
	}

	fmt.Fprintf(v, strings.Join(lines, "\n"))
	return nil
}
