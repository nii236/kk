package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// New initializes the singleton logger
func New(prod, debug bool) *logrus.Logger {
	log = logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	if prod {
		log.Infoln("Running in PRODUCTION mode")
	}
	if debug {
		log.Infoln("Running in DEBUG mode")
		log.Level = logrus.DebugLevel
	}
	return log
}

// Get returns the singleton instance of the logger
func Get() *logrus.Logger {
	if log == nil {
		panic("logger not initialized")
	}
	return log
}
