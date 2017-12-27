package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/olekukonko/tablewriter"
)

type Widget struct {
	Name     string
	Values   [][]string
	Selected int
	State    *k.State
}

func New(name string, initialState *k.State) *Widget {
	return &Widget{
		Name:  name,
		State: initialState,
	}
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

	v.Highlight = true
	v.Title = tw.State.UI.Table.Kind.String()
	v.SelBgColor = gocui.ColorCyan
	v.SelFgColor = gocui.ColorBlack
	v.Highlight = true

	if tw.State.UI.ActiveScreen == k.ScreenTable {
		_, err := g.SetCurrentView(k.ScreenTable.String())
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}

	}

	t := tablewriter.NewWriter(v)

	lines := [][]string{}
	switch tw.State.UI.Table.Kind {
	case k.KindPods:
		for _, pod := range tw.State.Entities.Pods.Pods.Items {
			lines = append(lines, k.PodLineHelper(pod))
		}
		if tw.State.Entities.Pods.Filter != "" {
			lines = filter(lines, func(s string) bool {
				if s == tw.State.Entities.Pods.Filter {
					return true
				}
				return false
			})
		}
		t.SetHeader(k.PodListHeaders)
		v.SetCursor(0, tw.State.Entities.Pods.Cursor)
		v.SetOrigin(0, tw.State.Entities.Pods.Cursor-maxY+10)
	case k.KindNamespaces:
		for _, ns := range tw.State.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, []string{ns.Name})
		}
		t.SetHeader(k.NamespaceListHeaders)
		v.SetCursor(0, tw.State.Entities.Namespaces.Cursor)
		v.SetOrigin(0, tw.State.Entities.Namespaces.Cursor-maxY+10)
	default:
		panic("Unsupported table kind: " + tw.State.UI.Table.Kind)
	}

	t.SetBorder(false)
	t.SetColumnSeparator("")
	t.AppendBulk(lines)
	// t.SetAlignment(tablewriter.ALIGN_CENTER)
	t.SetHeaderLine(false)
	t.Render()

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
