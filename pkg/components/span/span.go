package span

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type SpanText struct {
	Name     string
	Text     string
	Centered bool
	Pos      Position
}

type Position int

const (
	Top Position = iota + 1
	Bottom
)

func New(name, text string, centered bool, position Position) *SpanText {
	return &SpanText{
		Name:     name,
		Text:     text,
		Centered: centered,
		Pos:      position,
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
	text := st.Text
	if st.Centered {
		text = fmt.Sprintf("%[1]*s", -maxX, fmt.Sprintf("%[1]*s", (maxX+len(st.Text))/2, st.Text))
	}

	v, err := g.SetView(st.Name, x0, y0, x1, y1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Clear()
	v.Write([]byte(text))
	return nil
}
