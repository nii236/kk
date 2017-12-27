package k

import (
	"fmt"
	"os"
	"time"

	"github.com/jroimartin/gocui"
)

// DebugFileWriter returns a file for writing debug logs.
// Remember to close!
func DebugFileWriter(debugPath string) (*os.File, error) {
	_, err := os.Stat(debugPath)
	if os.IsNotExist(err) {
		f, err := os.Create(debugPath)
		if err != nil {
			return nil, err
		}
		err = f.Close()
		if err != nil {
			return nil, err
		}
	}
	f, err := os.OpenFile(debugPath, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return f, err
}

func Debugln(g *gocui.Gui, val interface{}) error {
	f, err := DebugFileWriter("/tmp/debug.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	v, err := g.View("Debug")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	t := time.Now()
	tf := t.Format("2006-01-02 15:04:05")
	fmt.Fprintf(v, "%s> %s\n", tf, val)
	fmt.Fprintf(f, "%s> %s\n", tf, val)
	return nil
}
