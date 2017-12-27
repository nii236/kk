package table

import (
	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
	"github.com/olekukonko/tablewriter"
)

// Widget is the table widget
type Widget struct {
	Name     string
	Values   [][]string
	Selected int
	State    *k.State
}

// New returns a new table widget
func New(name string, initialState *k.State) *Widget {
	return &Widget{
		Name:  name,
		State: initialState,
	}
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
	case k.KindTableDeployments:
		for _, deployment := range tw.State.Entities.Deployments.Deployments.Items {
			lines = append(lines, k.DeploymentLineHelper(deployment))
		}

		if tw.State.Entities.Deployments.Filter != "" {
			lines = filter(lines, func(s string) bool {
				if s == tw.State.Entities.Deployments.Filter {
					return true
				}
				return false
			})
		}
		t.SetHeader(k.DeploymentListHeaders)
		v.SetCursor(0, tw.State.Entities.Deployments.Cursor)
		v.SetOrigin(0, tw.State.Entities.Deployments.Cursor-maxY+10)
	case k.KindTablePods:
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
	case k.KindTableNamespaces:
		for _, ns := range tw.State.Entities.Namespaces.Namespaces.Items {
			lines = append(lines, k.NamespaceLineHelper(ns))
		}
		t.SetHeader(k.NamespaceListHeaders)
		t.SetColumnAlignment([]int{
			tablewriter.ALIGN_CENTER,
		})
		v.SetCursor(0, tw.State.Entities.Namespaces.Cursor)
		v.SetOrigin(0, tw.State.Entities.Namespaces.Cursor-maxY+10)
	default:
		panic("Unsupported table kind: " + tw.State.UI.Table.Kind)
	}

	t.SetBorder(false)
	t.SetColumnSeparator("")
	t.AppendBulk(lines)

	// t.SetAlignment(tablewriter.ALIGN_CENTER)
	t.SetColMinWidth(1, maxX-55)
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
