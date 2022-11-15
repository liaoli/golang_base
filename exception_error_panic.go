package main

import (
	"errors"
	"fmt"
)

/**
*@author: 廖理
*@date:2022/8/5
**/

//error 处理demo
func errorDemo() {
	var a int64 = -1
	r, err := errorTest(a)
	if err != nil {
		fmt.Printf("err:%v", err)
	} else {
		fmt.Printf("errorTest(%d):%v", a, r)
	}

}

//Go语言引入了一个关于错误处理的标准模式，即error接口，它是Go语言内建的接口类型
//type error interface {
//	Error() string
//}
func errorTest(a int64) (result int64, err error) {

	if a < 0 {
		err = errors.New("参数不能小于0")
		return
	}

	return a, err

}

//error返回的是一般性的错误，但是panic函数返回的是让程序崩溃的错误。
//panic 函数
//也就是当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等，这些运行时错误会引起painc异常，在一般情况下，我们不应通过调用panic函数来报告普通的错误，而应该只把它作为报告致命错误的一种方式。当某些不应该发生的场景发生时，我们就应该调用panic。
//一般而言，当panic异常发生时，程序会中断运行。随后，程序崩溃并输出日志信息。日志信息包括panic value和函数调用的堆栈跟踪信息。
//当然，如果直接调用内置的panic函数也会引发panic异常；panic函数接受任何值作为参数。
//下面给大家演示一下，直接调用panic函数，是否会导致程序的崩溃。

//所以，我们在实际的开发过程中并不会直接调用panic( )函数，但是当我们编程的程序遇到致命错误时，系统会自动调用该函数来终止整个程序的运行，也就是系统内置了panic函数。例如发现空指针或者数组越界时系统会自动调用空指针

func testPanic() {
	func1()
	func2()
	func3()
}

func func1() {
	fmt.Print("func1\n")
}

func func2() {
	panic("func2 panic")
}

func func3() {
	fmt.Print("func1\n")
}

//运行时panic异常一旦被引发就会导致程序崩溃。这当然不是我们愿意看到的，因为谁也不能保证程序不会发生任何运行时错误。
//Go语言为我们提供了专用于“拦截”运行时panic的内建函数——recover。它可以是当前的程序从运行时panic的状态中恢复并重新获得流程控制权。
//
//语法如下:
//func recover() interface{}
//注意：recover只有在defer调用的函数中有效。
//recover demo

//通过以上程序，我们发现虽然TestB()函数会导致整个应用程序崩溃，但是由于在改函数中调用了recover()函数，所以整个函数并没有崩溃。
//虽然程序没有崩溃，但是我们也没有看到任何的提示信息，那么怎样才能够看到相应的提示信息呢？
//可以直接打印recover()函数的返回结果
//recover()必须搭配defer使用。
//defer一定要在可能引发panic的语句之前定义。
func testRecover() {
	funcA()
	funcB()
	funcC()
}

func funcA() {
	fmt.Print("funcA\n")
}

func funcB() {
	defer func() {
		p := recover()
		fmt.Printf("recover result: %v\n", p)
	}()
	panic("func2 panic")
}

func funcC() {
	fmt.Print("funcC\n")
}
