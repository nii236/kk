package ui

import (
	"fmt"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	"github.com/nii236/k/pkg/components/modal"
	"github.com/nii236/k/pkg/components/span"
	"github.com/nii236/k/pkg/components/table"
	"github.com/nii236/k/pkg/k8s"
)

var stateChannel chan *k.State

var store *k.State

// App contains the TUI
type App struct {
	ClientSet *k8s.ClientSet
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
	stateChannel = make(chan *k.State)
	store = &k.State{}
}

// New returns a new instance of the TUI
func New(flags *k.ParsedFlags, clientSet *k8s.ClientSet) (*App, error) {
	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		return nil, err
	}

	app := &App{
		ClientSet: clientSet,
		Gui:       g,
	}

	app.ClientSet = clientSet
	app.Gui = g
	app.Gui.InputEsc = true
	app.Gui.SelFgColor = gocui.ColorGreen

	podList := table.New("Pods")

	// svcList := table.New("Services")
	titleSpan := span.New("Titlebar", "Kubectl TUI", true, span.Top)
	legendSpan := span.New("Legend", "^C: Exit ^R: Resource, ^N: Namespace ^L: Logs", true, span.Bottom)

	stateModal := modal.New("State", modal.Large, store)
	resourcesModal := modal.New("Resources", modal.Large, store)
	namespacesModal := modal.New("Namespaces", modal.Large, store)
	debugModal := modal.New("Debug", modal.Large, store)

	app.Gui.SetManager(namespacesModal, debugModal, stateModal, resourcesModal, podList, titleSpan, legendSpan)

	keys := []Key{
		Key{"", gocui.KeyCtrlC, exit},
		Key{"", gocui.KeyCtrlR, modal.Toggle(resourcesModal)},
		Key{"", gocui.KeyCtrlB, modal.Toggle(stateModal)},
		Key{"", gocui.KeyCtrlN, modal.Toggle(namespacesModal)},
		Key{"", gocui.KeyCtrlD, modal.Toggle(debugModal)},
		Key{"", gocui.KeyEsc, HideError},
		Key{"", 'L', ActionLoadMock},
		Key{"", gocui.KeyArrowUp, ActionPrev},
		Key{"", gocui.KeyArrowDown, ActionNext},
		// Key{"Resources", gocui.KeyArrowUp, modal.Prev(resources)},
		// Key{"Resources", gocui.KeyArrowDown, modal.Next(resources)},
		Key{"", gocui.KeyEnter, ActionSelectResource},
	}

	for _, key := range keys {
		if err := g.SetKeybinding(key.viewname, key.key, gocui.ModNone, key.handler); err != nil {
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
	// 			podLoader := Load(podList)
	// 			g.Update(podLoader)
	// 		}
	// 	}
	// }(t)

	go StateLoader(g)
	return app, nil
}

func exit(g *gocui.Gui, v *gocui.View) error {
	g.Close()
	os.Exit(0)
	return nil
}
