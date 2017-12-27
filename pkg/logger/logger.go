package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log wraps the logrus Logger
type Log struct {
	*logrus.Logger
}

var log *Log

// New initializes the singleton logger
func New(DebugToFile, Debug bool) {
	log = &Log{
		logrus.New(),
	}

	log.Formatter = &logrus.TextFormatter{ForceColors: true}

	if DebugToFile {
		h, err := NewLogrusFileHook("/tmp/debug.log", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		log.AddHook(h)
	}

	if Debug {
		log.Infoln("Running in DEBUG mode")
		log.SetLevel(logrus.DebugLevel)
	}
}

// Get returns the singleton instance of the logger
func Get() *Log {
	if log == nil {
		panic("logger not initialized")
	}
	return log
}
