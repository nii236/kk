package k

import (
	"github.com/nii236/k/pkg/logger"
	"github.com/pkg/errors"
)

// Infoln is a wrapper for logrus' Infoln
func Infoln(val ...interface{}) error {
	log := logger.Get()
	log.Infoln(val...)
	return nil
}

// Debugln is a wrapper for logrus' Debugln
func Debugln(val ...interface{}) error {
	log := logger.Get()
	log.Debugln(val...)
	return nil
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Errorln is a wrapper for logrus' Errorln
func Errorln(val ...interface{}) error {
	log := logger.Get()
	log.Errorln(val...)
	return nil
}
