package log

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

/**
*@author: 廖理
*@date:2022/12/7
**/

type FileLog struct {
	Level
	file *os.File
}

func NewFileLog(path string, l Level) *FileLog {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		panic("初始化log文件错误")
	}

	return &FileLog{
		Level: l,
		file:  f,
	}
}

func (f *FileLog) FinishLog() {
	f.file.Close()
}

func (f *FileLog) printLog(s string, v []interface{}, l Level) {
	t := time.Now().Format("2006-01-02 15:04:05.000")

	p, file, n, ok := runtime.Caller(2)
	if ok {

	}
	fun := runtime.FuncForPC(p)

	//fileFunc,nFunc:=fun.FileLine(p)
	//fmt.Printf("%s方法%s %d\n",fun.Name(),fileFunc,nFunc)
	content := fmt.Sprintf(s, v...)

	lDes := GetLevelDes(l)

	result := fmt.Sprintf("[%s],%s,%s %s:%s:%d\n", lDes, t, content, file, fun.Name(), n)

	f.file.WriteString(result)
}

func (f *FileLog) Error(s string, v ...interface{}) {
	f.printLog(s, v, error)
}

func (f *FileLog) Info(s string, v ...interface{}) {
	if f.Level > info {
		return
	}
	f.printLog(s, v, info)
}

func FLogDemo() {
	log := NewFileLog("./test.log", error)
	defer log.FinishLog()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			log.Info("logDemo %s", "好")
		} else {
			log.Error("logDemo %s", "报错了")
		}

		time.Sleep(time.Second)
	}
}
