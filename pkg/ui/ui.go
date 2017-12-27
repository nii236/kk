package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"k8s.io/api/core/v1"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/actions"
	"github.com/nii236/k/pkg/components/debug"
	"github.com/nii236/k/pkg/components/modal"
	"github.com/nii236/k/pkg/components/span"
	"github.com/nii236/k/pkg/components/table"
	"github.com/nii236/k/pkg/k"
	"github.com/nii236/k/pkg/k8s"
)

// App contains the TUI
type App struct {
	ClientSet k8s.ClientSet
	Gui       *gocui.Gui
}

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

func init() {
}

var store = &k.State{
	UI: &k.UIReducer{
		ActiveScreen: "Table",
		Table: &k.TableView{
			Kind:     "Pods",
			Selected: "",
			Filter:   "",
		},
		Modal: &k.ModalView{
			Cursor:   0,
			Kind:     k.KindNamespaces,
			Lines:    []string{},
			Selected: "",
		},
		State: &k.StateView{
			Cursor: 0,
		},
		Debug: &k.DebugView{
			Cursor: 0,
		},
	},
	Entities: &k.EntitiesReducer{
		Debug: &k.DebugEntities{
			Lines: []interface{}{},
		},
		Pods: &k.PodEntities{
			Cursor:         1,
			Loaded:         false,
			SendingRequest: false,
			Pods:           &v1.PodList{},
		},
		Errors: &k.ErrorEntities{
			Lines:        []string{},
			Acknowledged: true,
		},
		Namespaces: &k.NamespaceEntities{
			Cursor:         1,
			Loaded:         false,
			SendingRequest: false,
			Namespaces:     &v1.NamespaceList{},
		},
		Resources: &k.ResourceEntities{
			Resources: []string{k.KindNamespaces.String(), k.KindPods.String()},
		},
	},
}

// New returns a new instance of the TUI
func New(flags *k.ParsedFlags, clientSet *k8s.RealClientSet) (*App, error) {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return nil, err
	}

	app := &App{
		ClientSet: clientSet,
		Gui:       g,
	}
	app.ClientSet = clientSet
	if flags.TEST {
		app.ClientSet, err = k8s.NewMock(flags)
		if err != nil {
			panic(errors.Wrap(err, "Could not create mock client"))
		}
	}
	app.Gui = g
	app.Gui.InputEsc = true
	app.Gui.SelFgColor = gocui.ColorGreen

	tableView := table.New(k.ScreenTable.String(), store)
	modalView := modal.New(k.ScreenModal.String(), modal.Large, store)
	debugView := debug.New(k.ScreenDebug.String(), store)

	// svcList := table.New("Services")
	titleSpan := span.New("Titlebar", "Kubectl TUI", true, span.Top)
	legendSpan := span.New("Legend", "^c: Exit ^r: Resource, ^n: Filter ^f: Clear Filter L: Load Data", true, span.Bottom)

	app.Gui.SetManager(tableView, debugView, modalView, titleSpan, legendSpan)

	keys := []Key{
		Key{"", gocui.KeyCtrlC, exit},
		Key{"", gocui.KeyCtrlR, actions.ToggleResources(store)},
		Key{"", gocui.KeyCtrlN, actions.ToggleNamespaces(store)},
		Key{"", gocui.KeyCtrlD, actions.ToggleViewDebug(store)},
		Key{"", gocui.KeyEsc, actions.AcknowledgeErrors(store)},
		Key{"Table", 'd', actions.TableDelete(store, clientSet)},
		Key{"", 'D', actions.StateDump(store)},
		Key{"", 'L', actions.LoadManual(app.ClientSet, store)},
		Key{"", gocui.KeyArrowUp, actions.Prev(store)},
		Key{"", gocui.KeyPgup, actions.PageUp(store)},
		Key{"", gocui.KeyArrowDown, actions.Next(store)},
		Key{"", gocui.KeyPgdn, actions.PageDown(store)},
		Key{"Modal", gocui.KeyEnter, actions.HandleModalEnter(store)},
		Key{"Table", gocui.KeyEnter, actions.HandleTableEnter(store)},
		Key{"Table", gocui.KeyCtrlF, actions.TableClearFilter(store)},
		Key{"Table", gocui.KeyArrowUp, actions.TableCursorMoveUp(store)},
		Key{"Table", gocui.KeyArrowDown, actions.TableCursorMoveDown(store)},
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

					podLoader := actions.Load(g2, clientSet, store)
					g2.Update(podLoader)
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
