package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// LogrusFileHook is a file hook for logrus
type LogrusFileHook struct {
	file      *os.File
	flag      int
	chmod     os.FileMode
	formatter *logrus.TextFormatter
}

// NewLogrusFileHook returns a new logrus file hook
func NewLogrusFileHook(file string, flag int, chmod os.FileMode) (*LogrusFileHook, error) {
	logFile, err := os.OpenFile(file, flag, chmod)
	if os.IsNotExist(err) {
		logFile, err = os.Create(file)
		if err != nil {
			return nil, err
		}
	}
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook %v", err)
		return nil, err
	}

	return &LogrusFileHook{logFile, flag, chmod, &logrus.TextFormatter{
		ForceColors: true,
	}}, err
}

// Fire will execute when logrus is used
func (hook *LogrusFileHook) Fire(entry *logrus.Entry) error {

	plainformat, err := hook.formatter.Format(entry)
	line := string(plainformat)
	_, err = hook.file.WriteString(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to write file on filehook(entry.String)%v", err)
		return err
	}

	return nil
}

// Levels are the levels that the logrus file hook supports
func (hook *LogrusFileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}
