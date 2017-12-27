package k

import "github.com/nii236/k/pkg/logger"

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

// Errorln is a wrapper for logrus' Errorln
func Errorln(val ...interface{}) error {
	log := logger.Get()
	log.Errorln(val...)
	return nil
}
