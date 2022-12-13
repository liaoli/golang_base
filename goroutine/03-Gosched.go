package goroutine

import (
	"fmt"
	"runtime"
)

func GoroutienDemo3() {

	go func() {
		for {
			fmt.Println(" this is goroutine test")
		}
	}()

	for {
		runtime.Gosched() // 出让当前 cpu 时间片。
		fmt.Println(" this is main test")
	}
}

func GoroutienDemo() {
	//GoroutineDemo1()
	//GoroutienDemo2()
	//GoroutienDemo3()
	//GoroutienDemo4()
	//GoroutienDemo5()
	GoroutienDemo6()
}
