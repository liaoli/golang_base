package goroutine

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccccccccc")
	//return
	runtime.Goexit() // 退出当前go程。
	defer fmt.Println("ddddddddddddddddd")
}

func GoroutienDemo4() {

	go func() {
		defer fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbb")
	}()

	for {

	}
}
