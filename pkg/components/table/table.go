package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/olekukonko/tablewriter"
)

// Renderer contains the generic (!) methods needed
type Renderer interface {
	Lines(*k.State) [][]string
	Cursor(*k.State) int
	Headers(*k.State) []string
	Origin(*k.State, *gocui.View) (int, int)
}

// Widget is the table widget
type Widget struct {
	Name     string
	Values   [][]string
	Selected int
	State    *k.State
	Render   Renderer
}

// New returns a new table widget
func New(name string, initialState *k.State, renderer Renderer) *Widget {
	return &Widget{
		Render: renderer,
		Name:   name,
		State:  initialState,
	}
}

// Layout for the tablewidget
func (tw *Widget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	currentView := g.CurrentView()
	if currentView == nil {
		g.SetCurrentView(k.KindTablePods.String())
	}

	v, err := g.SetView(tw.Name, 0, 3, maxX-1, maxY-4)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	v.Highlight = true
	v.Title = tw.State.UI.Table.Kind.String()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	v.Highlight = true

	if tw.State.UI.ActiveScreen == k.ScreenTable {
		_, err := g.SetCurrentView(tw.State.UI.Table.Kind.String())
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		g.SetViewOnTop(tw.State.UI.Table.Kind.String())

	}

	cursor := tw.Render.Cursor(tw.State)
	lines := tw.Render.Lines(tw.State)
	headers := tw.Render.Headers(tw.State)
	ox, oy := tw.Render.Origin(tw.State, v)

	t := tablewriter.NewWriter(v)
	t.SetHeader(headers)
	t.SetBorder(false)
	t.SetColumnSeparator("")
	t.AppendBulk(lines)
	t.SetColMinWidth(1, maxX-55)
	t.SetHeaderLine(false)
	t.Render()

	v.SetCursor(0, cursor)
	v.SetOrigin(ox, oy)

	return nil
}

func filter(vs [][]string, f func(string) bool) [][]string {
	vsf := make([][]string, 0)
	for _, v := range vs {
		if f(v[0]) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
