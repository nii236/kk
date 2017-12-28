package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/actions"
	"github.com/nii236/k/pkg/components/debug"
	"github.com/nii236/k/pkg/components/modal"
	"github.com/nii236/k/pkg/components/span"
	"github.com/nii236/k/pkg/components/table"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
	"github.com/nii236/k/pkg/logger"
)

// App contains the TUI
type App struct {
	ClientSet k8s.ClientSet
	Gui       *gocui.Gui
}

// Key is a keybinding for the app
type Key struct {
	viewname string
	key      interface{}
	handler  func(*gocui.Gui, *gocui.View) error
}

// Run will execute the TUI
func (app *App) Run() error {
	if err := app.Gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
		app.Gui.Close()
		panic(err)
	}
	defer app.Gui.Close()

	return nil
}

var store = &k.State{
	UI: &k.UIReducer{
		ActiveScreen: k.ScreenTable,
		Table: &k.TableView{
			Kind: k.KindTablePods,
		},
		Modal: &k.ModalView{
			Cursor: 0,
			Kind:   k.KindModalNamespaces,
			Lines:  []string{},
			Size:   k.ModalSizeLarge,
		},
	},
	Entities: &k.EntitiesReducer{
		Deployments: &k.DeploymentEntities{
			Cursor:         0,
			Size:           0,
			SendingRequest: false,
			Deployments:    &appsv1.DeploymentList{},
		},
		Pods: &k.PodEntities{
			Cursor:         0,
			Size:           0,
			SendingRequest: false,
			Pods:           &v1.PodList{},
		},
		Namespaces: &k.NamespaceEntities{
			Cursor:         1,
			Size:           0,
			SendingRequest: false,
			Namespaces:     &v1.NamespaceList{},
		},
		Resources: &k.ResourceEntities{
			Resources: []string{k.KindTableNamespaces.String(), k.KindTablePods.String(), k.KindTableDeployments.String()},
		},
	},
}

// New returns a new instance of the TUI
func New(flags *k.ParsedFlags, clientSet k8s.ClientSet) (*App, error) {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return nil, err
	}

	log := logger.Get()
	log.AddHook(logger.NewGocuiHook(g))

	app := &App{
		ClientSet: clientSet,
		Gui:       g,
	}
	app.ClientSet = clientSet
	if flags.Test {
		app.ClientSet, err = k8s.NewMock(flags)
		if err != nil {
			panic(errors.Wrap(err, "Could not create mock client"))
		}
	}
	app.Gui = g
	app.Gui.InputEsc = true
	app.Gui.SelFgColor = gocui.ColorGreen
	// tableView := table.New(k.ScreenTable.String(), store)
	deploymentsView := table.New(k.KindTableDeployments.String(), store, table.NewDeploymentRenderer())
	namespacesView := table.New(k.KindTableNamespaces.String(), store, table.NewNamespaceRenderer())
	podView := table.New(k.KindTablePods.String(), store, table.NewPodRenderer())
	modalView := modal.New(k.ScreenModal.String(), store)
	debugView := debug.New(k.ScreenDebug.String(), store)

	// svcList := table.New("Services")
	titleSpan := span.New("Title", "Kubectl TUI", true, span.Top, store)
	legendSpan := span.New("Legend", "", true, span.Bottom, store)

	app.Gui.SetManager(debugView, podView, deploymentsView, namespacesView, modalView, titleSpan, legendSpan)
	keys := []Key{
		Key{"", gocui.KeyCtrlC, exit},
		Key{"", gocui.KeyCtrlR, actions.ToggleResources(store)},
		Key{"", gocui.KeyCtrlN, actions.ToggleNamespaces(store)},
		Key{"", gocui.KeyCtrlD, actions.ToggleViewDebug(store)},
		Key{"Debug", gocui.KeyEsc, actions.HandleDebugEsc(store)},
		Key{"Modal", gocui.KeyEsc, actions.HandleModalEsc(store)},
		Key{"", 'd', actions.HandleTableDelete(store, clientSet)},
		Key{"", 'D', actions.StateDump(store)},
		Key{"", 'L', actions.LoadManual(app.ClientSet, store)},
		Key{"Modal", gocui.KeyArrowUp, actions.ModalCursorMove(store, -1)},
		Key{"Modal", gocui.KeyPgup, actions.ModalCursorMove(store, -5)},
		Key{"Modal", gocui.KeyArrowDown, actions.ModalCursorMove(store, 1)},
		Key{"Modal", gocui.KeyPgdn, actions.ModalCursorMove(store, 5)},
		Key{"Modal", gocui.KeyEnter, actions.HandleModalEnter(store, clientSet)},
		Key{k.KindTablePods.String(), gocui.KeyEnter, actions.HandleTableEnter(store, clientSet)},
		Key{k.KindTablePods.String(), gocui.KeyCtrlF, actions.TableClearFilter(store)},
		Key{k.KindTablePods.String(), gocui.KeyArrowUp, actions.TableCursorMove(store, -1)},
		Key{k.KindTablePods.String(), gocui.KeyArrowDown, actions.TableCursorMove(store, 1)},
		Key{k.KindTablePods.String(), gocui.KeyPgup, actions.TableCursorMove(store, -5)},
		Key{k.KindTablePods.String(), gocui.KeyPgdn, actions.TableCursorMove(store, 5)},

		Key{k.KindTableDeployments.String(), gocui.KeyArrowUp, actions.TableCursorMove(store, -1)},
		Key{k.KindTableDeployments.String(), gocui.KeyArrowDown, actions.TableCursorMove(store, 1)},
		Key{k.KindTableDeployments.String(), gocui.KeyPgup, actions.TableCursorMove(store, -5)},
		Key{k.KindTableDeployments.String(), gocui.KeyPgdn, actions.TableCursorMove(store, 5)},

		Key{k.KindTableNamespaces.String(), gocui.KeyArrowUp, actions.TableCursorMove(store, -1)},
		Key{k.KindTableNamespaces.String(), gocui.KeyArrowDown, actions.TableCursorMove(store, 1)},
		Key{k.KindTableNamespaces.String(), gocui.KeyPgup, actions.TableCursorMove(store, -5)},
		Key{k.KindTableNamespaces.String(), gocui.KeyPgdn, actions.TableCursorMove(store, 5)},
	}

	for _, key := range keys {
		if err := app.Gui.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	if flags.AutoRefresh {
		t := time.NewTicker(time.Duration(flags.RefreshInterval) * time.Second)
		go func(g2 *gocui.Gui, t *time.Ticker, clientSet k8s.ClientSet, store *k.State) {
			for {
				select {
				case <-t.C:

					podLoader := actions.LoadAuto(clientSet, store)
					tableView, err := g2.View(k.KindTablePods.String())
					if err != nil {
						return
					}
					podLoader(g2, tableView)
				}
			}
		}(g, t, clientSet, store)

	}

	return app, nil
}

func exit(g *gocui.Gui, v *gocui.View) error {
	g.Close()
	os.Exit(0)
	return nil
}
