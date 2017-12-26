package debug

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/common"
)

type Widget struct {
	Name string
}

//ss
func New(name string) *Widget {
	return &Widget{
		Name: name,
	}
}

// Layout for the tablewidget
func (st *Widget) Layout(g *gocui.Gui) error {
	w, h := g.Size()
	v, err := g.SetView(k.ScreenDebug.String(), 0, 3, w-1, h-4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	store, err := common.JSONToState(g)
	if err != nil {
		fmt.Fprint(v, err)
		return nil
	}
	if store.UI.ActiveScreen != k.ScreenDebug {
		g.SetViewOnBottom(k.ScreenDebug.String())
		return nil
	}

	g.SetViewOnTop(k.ScreenDebug.String())
	g.SetCurrentView(k.ScreenDebug.String())
	v.Title = "Debug"
	v.Clear()
	v.Autoscroll = true
	lines := []string{}
	for _, el := range store.Entities.Debug.Lines {
		lines = append(lines, el.(string))
	}

	fmt.Fprintf(v, strings.Join(lines, "\n"))
	return nil
}
