package logger

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

// Log wraps the logrus Logger
type Log struct {
	*logrus.Logger
}

var log *Log

// New initializes the singleton logger
func New(prod, debug bool) {
	log = &Log{
		logrus.New(),
	}

	log.Out = ioutil.Discard
	log.Formatter = &logrus.TextFormatter{ForceColors: true}

	if debug {
		log.Level = logrus.DebugLevel
		h, err := NewLogrusFileHook("/tmp/debug.log", os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}
		log.AddHook(h)
		log.Infoln("Running in DEBUG mode")
	}
	if prod {
		log.Infoln("Running in PRODUCTION mode")
	}
}

// Get returns the singleton instance of the logger
func Get() *Log {
	if log == nil {
		panic("logger not initialized")
	}
	return log
}
