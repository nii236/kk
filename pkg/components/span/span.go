package span

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

type SpanText struct {
	Name     string
	Text     string
	Centered bool
	Pos      Position
	State    *k.State
}

type Position int

const (
	Top Position = iota + 1
	Bottom
)

func New(name, text string, centered bool, position Position, store *k.State) *SpanText {
	return &SpanText{
		Name:     name,
		Text:     text,
		Centered: centered,
		Pos:      position,
		State:    store,
	}
}

func (st *SpanText) SetVal(val string) {
	st.Text = val
}

func (st *SpanText) Val() string {
	return st.Text
}

// Layout for the tablewidget
func (st *SpanText) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	heightOffset := 0
	if st.Pos == Bottom {
		heightOffset = maxY - 3
	}
	x0 := 0
	x1 := maxX - 1
	y0 := heightOffset
	y1 := heightOffset + 2

	v, err := g.SetView(st.Name, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	if st.Name == "Legend" {
		text := ""
		switch st.State.UI.ActiveScreen {
		case k.ScreenDebug:
			text = getDebugLegend()
		case k.ScreenModal:
			text = getModalLegend(st.State.UI.Modal.Kind)
		case k.ScreenTable:
			text = getTableLegend(st.State.UI.Table.Kind)
		default:
			k.Errorln("Unsupported Screen:", st.State.UI.ActiveScreen)

		}
		centeredText := fmt.Sprintf("%[1]*s", -maxX, fmt.Sprintf("%[1]*s", (maxX+len(text))/2, text))
		v.Write([]byte(centeredText))
		return nil
	}

	if st.Name == "Title" {
		text := st.Text
		if st.Centered {
			text = fmt.Sprintf("%[1]*s", -maxX, fmt.Sprintf("%[1]*s", (maxX+len(st.Text))/2, st.Text))
		}
		v.Write([]byte(text))
		return nil
	}

	return nil
}

var Legend = map[string][]string{}

func getDebugLegend() string {
	result := "Esc: Back"
	return result
}

func getModalLegend(kind k.ModalKind) string {
	result := ""
	switch kind {
	case k.KindModalContainerLogs:
		result = "^c: Exit ^d: Logs L: Load Data D: Dump State Esc: Back"
	case k.KindModalNamespaces:
		result = "^c: Exit ^d: Logs L: Load Data D: Dump State Esc: Back"
	case k.KindModalResources:
		result = "^c: Exit ^d: Logs L: Load Data D: Dump State Esc: Back"
	case k.KindModalSelectContainer:
		result = "^c: Exit ^d: Logs L: Load Data D: Dump State Esc: Back"
	default:
		k.Errorln("Unsupported modal kind:", kind)
		result = "Unsupported modal kind"
	}
	return result
}

func getTableLegend(kind k.TableKind) string {
	result := ""
	switch kind {
	case k.KindTableDeployments:
		result = "^c: Exit ^r: Resource, ^n: Filter ^f: Clear Filter ^d: Logs L: Load Data, d: Delete Pod D: Dump State"
	case k.KindTableNamespaces:
		result = "^c: Exit ^r: Resource, ^n: Filter ^f: Clear Filter ^d: Logs L: Load Data, d: Delete Pod D: Dump State"
	case k.KindTablePods:
		result = "^c: Exit ^r: Resource, ^n: Filter ^f: Clear Filter ^d: Logs L: Load Data, d: Delete Pod D: Dump State"
	default:
		k.Errorln("Unsupported table kind:", kind)
		result = "Unsupported table kind"
	}
	return result
}
