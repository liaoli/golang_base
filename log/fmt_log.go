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

type Level int8

const (
	info Level = iota
	error
)

type FmtLogger struct {
	level Level
}

func (f *FmtLogger) Error(s string, v ...interface{}) {
	f.printLog(s, v, error)
}

func (f *FmtLogger) Info(s string, v ...interface{}) {
	if f.level > info {
		return
	}
	f.printLog(s, v, info)
}

func (f FmtLogger) GetLevelDes(l Level) string {
	switch l {
	case info:
		return "info"
	case error:
		return "error"
	}
	return ""
}

func (f *FmtLogger) printLog(s string, v []interface{}, l Level) {
	t := time.Now().Format("2006-01-02 15:04:05.000")

	p, file, n, ok := runtime.Caller(2)
	if ok {

	}
	fun := runtime.FuncForPC(p)

	//fileFunc,nFunc:=fun.FileLine(p)
	//fmt.Printf("%s方法%s %d\n",fun.Name(),fileFunc,nFunc)
	content := fmt.Sprintf(s, v...)

	lDes := f.GetLevelDes(l)

	fmt.Printf("[%s],%s,%s %s:%s:%d\n", lDes, t, content, file, fun.Name(), n)
}

func LogDemo() {
	log := &FmtLogger{
		level: error,
	}
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			log.Info("logDemo %s", "好")
		} else {
			log.Error("logDemo %s", "报错了")
		}

		time.Sleep(time.Second)
	}

}

func GetUseFunctionFileAndLineNumber() *FmtLogger {
	log := &FmtLogger{}
	p, _, _, ok := runtime.Caller(1)
	if ok {

	}
	fun := runtime.FuncForPC(p)

	fileFunc, nFunc := fun.FileLine(p)
	fmt.Printf("%s方法%s %d\n", fun.Name(), fileFunc, nFunc)
	return log
}
