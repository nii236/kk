package k

import (
	"fmt"
	"os"
	"time"
)

func Debugln(val interface{}) {
	f, err := os.OpenFile("/tmp/debug.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil && err == os.ErrNotExist {
		f, err = os.Create("/tmp/debug.log")
		if err != nil {
			panic(err)
		}
	}

	defer f.Close()
	t := time.Now()
	tf := t.Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("%s>%s\n", tf, val))

}
