package k

import (
	"github.com/jroimartin/gocui"
)

// TableView is the state for the table component
type TableView struct {
	Kind TableKind
}

const (
	// KindTableDeployments is the table of type Deployments
	KindTableDeployments TableKind = "DeploymentsTable"
	// KindTablePods is the table of type Pods
	KindTablePods TableKind = "PodsTable"
	// KindTableNamespaces is the table of type Namespaces
	KindTableNamespaces TableKind = "NamespacesTable"
)

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
