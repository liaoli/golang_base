package myhttp

import (
	"fmt"
	"runtime"
	"time"
)

func aaa() {
	for {
		time.Sleep(time.Millisecond * 200)
		fmt.Println("-----------\n")
	}
}
func gcDemo() {

	go func() {
		fmt.Println("------------1\n")
		go aaa()
		fmt.Println("------------2\n")
		return
	}()

	for {
		runtime.GC()
	}
}

func HttpDemo() {
	//gcDemo()
	//HttpServerDemo2()
	//HttpClientDemo1()
	//HttpServerDemo4()
	HttpServerDemo5()
}
