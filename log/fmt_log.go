package log

import (
	"fmt"
	"runtime"
	"time"
)

/**
*@author: 廖理
*@date:2022/12/5
**/

type FmtLogger struct {
}

func (f *FmtLogger) Info(s string, v ...interface{}) {
	t := time.Now().Format("2006-01-02 15:04:05.000")

	p, file, n, ok := runtime.Caller(0)
	if ok {

	}
	fun := runtime.FuncForPC(p).Name()
	content := fmt.Sprintf(s, v...)

	fmt.Printf("%s,[%s],%s %s:%s:%d\n", t, "info", content, file, fun, n)
}

func LogDemo() {
	log := &FmtLogger{}

	for {
		log.Info("logDemo %s", "好")
		time.Sleep(time.Second)
	}

}
