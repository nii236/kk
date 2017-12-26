package table

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/common"
	"github.com/olekukonko/tablewriter"
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
	v.Clear()

	store, err := common.JSONToState(g)
	if err != nil {
		fmt.Fprint(v, err)
		return nil
	}
	v.Highlight = true
	v.Title = store.UI.Table.Kind.String()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	v.Highlight = true

	if store.UI.ActiveScreen == k.ScreenTable {
		_, err := g.SetCurrentView(k.ScreenTable.String())
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}

	}

	lines := store.UI.Table.Lines
	t := tablewriter.NewWriter(v)
	t.SetBorder(false)
	t.SetColumnSeparator("")
	t.SetHeader(store.UI.Table.Headers)
	t.AppendBulk(lines)
	t.SetAlignment(tablewriter.ALIGN_CENTER)
	t.SetHeaderLine(false)
	t.Render()

	v.SetCursor(0, store.UI.Table.Cursor)
	v.SetOrigin(0, store.UI.Table.Cursor-maxY+10)

	return nil
}
