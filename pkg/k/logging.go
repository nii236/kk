package k

import "github.com/nii236/k/pkg/logger"

func Infoln(val ...interface{}) error {
	log := logger.Get()
	log.Infoln(val...)
	return nil
}
func Debugln(val ...interface{}) error {
	log := logger.Get()
	log.Debugln(val...)
	return nil
}

func Errorln(val ...interface{}) error {
	log := logger.Get()
	log.Errorln(val...)
	return nil
}
