package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

// Useful to debug Pody (display with CTRL+D)
// func debug(g *gocui.Gui, output interface{}) {
// 	v, err := g.View("Debug")
// 	if err == nil {
// 		t := time.Now()
// 		tf := t.Format("2006-01-02 15:04:05")
// 		output = tf + " > " + output.(string)
// 		fmt.Fprintln(v, output)
// 	}
// }

func displayError(g *gocui.Gui, e error) error {
	lMaxX, lMaxY := g.Size()
	minX := lMaxX / 6
	minY := lMaxY / 6
	maxX := 5 * (lMaxX / 6)
	maxY := 5 * (lMaxY / 6)

	if v, err := g.SetView("errors", minX, minY, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Settings
		v.Title = " ERROR "
		v.Frame = true
		v.Wrap = true
		v.Autoscroll = true
		v.BgColor = gocui.ColorRed
		v.FgColor = gocui.ColorWhite

		// Content
		v.Clear()
		fmt.Fprintln(v, e.Error())

		// Send to forground
		g.SetCurrentView(v.Name())
	}

	return nil
}
