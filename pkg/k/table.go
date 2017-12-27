package k

import (
	"github.com/pkg/errors"

	"github.com/jroimartin/gocui"
)

// TableView is the state for the table component
type TableView struct {
	Kind TableKind
}

const (
	// KindTableDeployments is the table of type Deployments
	KindTableDeployments TableKind = "Deployments"
	// KindTablePods is the table of type Pods
	KindTablePods TableKind = "Pods"
	// KindTableNamespaces is the table of type Namespaces
	KindTableNamespaces TableKind = "Namespaces"
)

// SelectResource updates the table view state with a new selected resource to display
func (v *TableView) SelectResource(g1 *gocui.Gui, resource string) {
	Debugln("UI TableView: SelectResource")
	g1.Update(
		func(g *gocui.Gui) error {
			switch resource {
			case KindTableNamespaces.String():
				Debugln("SelectResource Namespace")
				v.SetKind(g, KindTableNamespaces)
			case KindTablePods.String():
				Debugln("SelectResource Pod")
				v.SetKind(g, KindTablePods)
			case KindTableDeployments.String():
				Debugln("SelectResource Deployments")
				v.SetKind(g, KindTableDeployments)
			default:
				Errorln(errors.New("Unsupported resource: " + resource))
			}
			return nil

		},
	)
}

// SetKind updates the table view state with a new table kind
func (v *TableView) SetKind(g1 *gocui.Gui, kind TableKind) {
	Debugln("UI TableView: SetKind")
	g1.Update(
		func(g *gocui.Gui) error {
			v.Kind = kind
			return nil
		},
	)
}
