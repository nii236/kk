package state

import (
	"encoding/json"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k"
	v1 "k8s.io/api/core/v1"
)

type Widget struct {
	Name     string
	Size     SizeEnum
	Selected int
	Visible  bool
	State    *k.State
}

type SizeEnum int

const (
	Small SizeEnum = iota + 1
	Medium
	Large
)

func New(name string) *Widget {
	initialState := &k.State{
		UI: &k.UIReducer{
			ActiveScreen: "Table",
			Table: &k.TableView{
				Cursor:   0,
				Lines:    [][]string{},
				Selected: "",
			},
			Modal: &k.ModalView{
				Cursor:   0,
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
				Loaded:         false,
				SendingRequest: false,
				Pods:           &v1.PodList{},
			},
		},
	}
	return &Widget{
		Name:    name,
		Visible: false,
		State:   initialState,
	}
}

// Layout for the tablewidget
func (st *Widget) Layout(g *gocui.Gui) error {
	w, h := g.Size()
	v, err := g.SetView(k.ScreenState.String(), 0, 3, w-1, h-4) // x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	// v.Highlight = true
	v.Frame = true
	if st.State.UI.ActiveScreen != k.ScreenState {
		v, err := g.SetView(k.ScreenState.String(), 0, 0, 1, 1)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
	}

	g.SetViewOnTop(k.ScreenState.String())
	g.SetCurrentView(k.ScreenState.String())
	// v.Title = k.ScreenState.String()
	v.Title = "State"
	v.Clear()
	v.SetOrigin(0, st.State.UI.State.Cursor)
	v.Highlight = true
	v.SetCursor(0, st.State.UI.State.Cursor)
	enc := json.NewEncoder(v)
	enc.SetIndent("", "    ")
	return enc.Encode(st.State)
}

func Toggle(w *Widget) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		g.SetCurrentView(w.Name)
		w.Visible = !w.Visible
		return nil
	}
}
