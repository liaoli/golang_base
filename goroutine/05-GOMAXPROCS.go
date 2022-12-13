package goroutine

import (
	"fmt"
	"runtime"
)

func GoroutienDemo5() {

	fmt.Println(runtime.GOROOT())

	n := runtime.GOMAXPROCS(0) //将cpu设置为 单核
	fmt.Println("n = ", n)

	//n = runtime.GOMAXPROCS(2)  //将cpu设置为 双核
	//fmt.Println("n = ", n)
	//
	//n = runtime.GOMAXPROCS(2)  //将cpu设置为 双核
	//fmt.Println("n = ", n)

	//会卡死
	//go func() {
	//	for {
	//		fmt.Print(0) // 子go 程
	//	}
	//}()
	//
	//for {
	//	fmt.Print(1) // 主 go 程
	//}
}
