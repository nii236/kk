package ui

import (
	"github.com/jroimartin/gocui"
)

func ActionLoadMock(g *gocui.Gui, _ *gocui.View) error {
	Debug(g, "Run ActionLoadMock\n")
	LoadPods(g)
	LoadNamespaces(g)
	LoadResources(g)
	return nil
}

func ActionPrev(g *gocui.Gui, _ *gocui.View) error {
	v, _ := g.View("Pods")
	cx, cy := v.Cursor()
	v.SetCursor(cx, cy-1)
	// Debug(g, "HI")
	return nil
}

func ActionNext(g *gocui.Gui, _ *gocui.View) error {
	v, _ := g.View("Pods")
	cx, cy := v.Cursor()

	v.SetCursor(cx, cy+1)
	return nil
}

func ActionSelectResource(g *gocui.Gui, _ *gocui.View) error {
	Debug(g, "heyo")
	return nil
}

func HideError(g *gocui.Gui, _ *gocui.View) error {
	hideErrorPopup(g)
	return nil
}
