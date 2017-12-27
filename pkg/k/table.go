package k

import "github.com/jroimartin/gocui"

type TableView struct {
	Selected string
	Filter   string
	Kind     TableKind
}

const (
	KindTableDeployments TableKind = "Deployments"
	KindTablePods        TableKind = "Pods"
	KindTableNamespaces  TableKind = "Namespaces"
)

func (ur *TableView) SelectResource(g1 *gocui.Gui, resource string) {
	g1.Update(
		func(g *gocui.Gui) error {
			switch resource {
			case KindTableNamespaces.String():
				Debugln("SelectResource Namespace")
				ur.SetKind(g, KindTableNamespaces)
			case KindTablePods.String():
				Debugln("SelectResource Pod")
				ur.SetKind(g, KindTablePods)
			case KindTableDeployments.String():
				Debugln("SelectResource Deployments")
				ur.SetKind(g, KindTableDeployments)
			default:
				Errorln("Unsupported resource: " + resource)
			}
			return nil

		},
	)
}

func (p *TableView) SetKind(g1 *gocui.Gui, kind TableKind) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Kind = kind
			return nil
		},
	)
}

func (p *TableView) ClearFilter(g1 *gocui.Gui) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Filter = ""
			return nil
		},
	)
}

func (p *TableView) SetFilter(g1 *gocui.Gui, filter string) {
	g1.Update(
		func(g *gocui.Gui) error {
			p.Filter = filter
			return nil
		},
	)
}
