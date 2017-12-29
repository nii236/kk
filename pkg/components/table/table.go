package table

import (
	"github.com/jroimartin/gocui"
)

type Widget struct {
	Name     string
	Values   [][]string
	Selected int
}

func New(name string) *Widget {
	return &Widget{Name: name}
}

func (tw *Widget) Delta(delta int) {
	tw.Selected = tw.Selected + delta

	if tw.Selected < 0 {
		tw.Selected = 0
	}
	maxLength := len(tw.Values)
	if tw.Selected > maxLength {
		tw.Selected = maxLength
	}

}

func (tw *Widget) Val() [][]string {
	return [][]string{}
}

func (tw *Widget) SetVal(values [][]string) error {
	tw.Values = values
	return nil
}

// Layout for the tablewidget
func (tw *Widget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(tw.Name, 0, 3, maxX-1, maxY-4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	g.SetCurrentView(tw.Name)
	v.Highlight = true
	v.Title = tw.Name

	return nil
}
