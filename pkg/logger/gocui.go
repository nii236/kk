package logger

import (
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

// GocuiHook is the gocui logrus hook
type GocuiHook struct {
	g         *gocui.Gui
	formatter *logrus.TextFormatter
}

// NewGocuiHook returns a logrus hook that writes to the app's debug
func NewGocuiHook(g *gocui.Gui) *GocuiHook {
	return &GocuiHook{g, &logrus.TextFormatter{
		ForceColors: true,
	}}
}

// Fire event
func (hook *GocuiHook) Fire(entry *logrus.Entry) error {
	v, err := hook.g.View("Debug")
	if err != nil {
		return err
	}
	line, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}

	_, err = v.Write(line)
	if err != nil {
		return err
	}
	return nil
}

func (hook *GocuiHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
