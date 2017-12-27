package modal

import (
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/nii236/k/pkg/k"
)

// Widget represents a modal widget
type Widget struct {
	State *k.State
}

// New returns a new modal widget
func New(name string, initialState *k.State) *Widget {
	return &Widget{
		State: initialState,
	}
}

// Layout for the modal widget
func (st *Widget) Layout(g *gocui.Gui) error {

	w, h := g.Size()
	modalWidth := 0
	modalHeight := 0
	switch st.State.UI.Modal.Size {
	case k.ModalSizeSmall:
		modalWidth = w * 1 / 8
		modalHeight = h * 1 / 8
	case k.ModalSizeMedium:
		modalWidth = w * 1 / 4
		modalHeight = h * 1 / 4
	case k.ModalSizeLarge:
		modalWidth = w * 1 / 2
		modalHeight = h * 1 / 2
	default:
		modalWidth = w
		modalHeight = h
	}

	x0 := w/2 - modalWidth/2
	x1 := w/2 + modalWidth/2
	y0 := h/2 - modalHeight/2
	y1 := h/2 + modalHeight/2
	v, err := g.SetView(k.ScreenModal.String(), x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	// Extra large modals to be same size as table
	if st.State.UI.Modal.Size == k.ModalSizeExtraLarge {
		v, err = g.SetView(k.ScreenModal.String(), 0, 3, w-1, h-4)
		if err != nil && err != gocui.ErrUnknownView {
			return err
		}
	}

	if st.State.UI.ActiveScreen != k.ScreenModal {
		g.SetViewOnBottom(k.ScreenModal.String())
		return nil
	}

	g.SetCurrentView(k.ScreenModal.String())
	g.SetViewOnTop(k.ScreenModal.String())
	v.Clear()
	_, vy := v.Size()
	v.Highlight = true
	v.Title = st.State.UI.Modal.Title
	v.SetCursor(0, st.State.UI.Modal.Cursor)

	if st.State.UI.Modal.Cursor < vy {
		v.SetOrigin(0, 0)
	} else {
		v.SetOrigin(0, st.State.UI.Modal.Cursor-vy+1)
	}
	lines := st.State.UI.Modal.Lines
	v.Write([]byte(strings.Join(lines, "\n")))
	return nil
}
