package ui

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/actions"
	"github.com/nii236/k/pkg/components/debug"
	"github.com/nii236/k/pkg/components/modal"
	"github.com/nii236/k/pkg/components/span"
	"github.com/nii236/k/pkg/components/state"
	"github.com/nii236/k/pkg/components/table"
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

	tableView := table.New(k.ScreenTable.String())
	modalView := modal.New(k.ScreenModal.String(), modal.Large)
	store := state.New(k.ScreenState.String())
	debugView := debug.New(k.ScreenDebug.String())

	// svcList := table.New("Services")
	titleSpan := span.New("Titlebar", "Kubectl TUI", true, span.Top)
	legendSpan := span.New("Legend", "^C: Exit ^R: Resource, ^N: Namespace ^L: Logs", true, span.Bottom)

	app.Gui.SetManager(store, tableView, debugView, modalView, titleSpan, legendSpan)

	keys := []Key{
		Key{"", gocui.KeyCtrlC, exit},
		Key{"", gocui.KeyCtrlR, actions.ToggleResources(store)},
		Key{"", gocui.KeyCtrlN, actions.ToggleNamespaces(store)},
		Key{"", gocui.KeyCtrlD, actions.ToggleViewDebug(store)},
		Key{"", gocui.KeyCtrlB, actions.ToggleState(store)},
		Key{"", gocui.KeyEsc, actions.AcknowledgeErrors(store)},
		Key{"", 'L', actions.LoadMock(app.ClientSet, store)},
		Key{"", gocui.KeyArrowUp, actions.Prev(store)},
		Key{"", gocui.KeyPgup, actions.PageUp(store)},
		Key{"", gocui.KeyArrowDown, actions.Next(store)},
		Key{"", gocui.KeyPgdn, actions.PageDown(store)},
		Key{"", gocui.KeyEnter, actions.SelectResource(store)},
	}

	for _, key := range keys {
		if err := app.Gui.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	// t := time.NewTicker(1 * time.Second)
	// go func(t *time.Ticker) {
	// 	for {
	// 		select {
	// 		case <-t.C:
	// 			data := &v1.PodList{}
	// 			podLoader := LoadPods(g, store, data)
	// 			g.Update(podLoader)
	// 		}
	// 	}
	// }(t)

	return app, nil
}

func exit(g *gocui.Gui, v *gocui.View) error {
	g.Close()
	os.Exit(0)
	return nil
}
