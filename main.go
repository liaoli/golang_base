package main

import (
	"fmt"
)

var a int64

func main() {

	//scope()
	//
	//anonymousFunctionDemo()
	//
	//deferDemo1()

	//deferDemo2()

	//deferDemo3()

	//oopDemo()

	//referencePointer()
	//testPeptide()

	//testEmptyInterface()
	//errorDemo()
	//testPanic()

	//testRecover()
	//stringDemo()
	//deleteSliceItem()

	//r:=rand.Perm(10) //获取随机数的工具类
	//
	//fmt.Print(r)

	//structDemo()
	//FileDemo()
	//goroutinesDemo()
	//lockDemo()
	//syncSysMap()
	syncMapDemo()
}

//作用域的问题
func scope() {
	//作用域的问题
	//a := 9
	a = 9
	Test()
	fmt.Printf("main a =%d\n", a)
}

//匿名函数 例子
func anonymousFunctionDemo() {
	//匿名函数

	//在这里，有一件非常有意思的事情，就是在匿名函数中可以直接访问main( )函数中定义的局部变量，
	//？并且在匿名函数中对变量的值进行了修改，最终会影响到整个main( )函数中定义的变量的值

	m := 5

	f := func() {
		m = 100
		fmt.Printf("f m = %d\n", m)
	}

	f()

	fmt.Printf("main m = %d\n", m)

	var tf TypeFunc
	tf = f

	tf()

	//定义匿名函数时，直接调用
	func() {
		fmt.Printf("匿名函数直接调用\n")
	}()
}

func Test() {
	a = 5
	a += 10
	fmt.Printf("test a =%d\n", a)
}

//TypeFunc 自定义一个函数类型 此类型函数可以接受没有参数没有返回值的任意函数变量
type TypeFunc func()

func closer() func() int64 {
	var x int64
	return func() int64 {
		x++
		return x
	}
}

//defer demo 1
func deferDemo1() {
	//closer( )函数时指定了返回的类型是一个匿名函数，并且该匿名函数返回的类型是整型。
	//closer( )函数中定义了一个匿名函数，并且将整个匿名函数返回，匿名函数返回的是整型。
	//在main( )函数中定义了一个变量f,该变量的类型是匿名函数，f( )表示调用执行匿名函数。
	//最终执行完成后发现，实现了数字的累加。
	//因为匿名函数(闭包)，有一个很重要的特点:
	//它不关心这些捕获了的变量和常量是否已经超出了作用域，所以只有闭包还在使用它，这些变量就还会存在。

	c := closer()
	fmt.Printf("闭包1次：%d\n", c())
	fmt.Printf("闭包2次：%d\n", c())
	fmt.Printf("闭包3次：%d\n", c())

	//defer 如果一个函数中有多个defer语句，它们会以LIFO（后进先出）的顺序执行。
	defer fmt.Printf("defer 1\n")
	defer fmt.Printf("defer 2\n")
	defer func() {
		var x, b float64
		x = 3
		b = 0.0
		y := x / b
		defer fmt.Printf("b/x= %f \n", y)
	}()
	defer fmt.Printf("defer 3\n")
}

//deferDemo2 defer 执行的时间
func deferDemo2() {
	f := f1
	r := f()
	fmt.Printf("r = %d\n", r) //r=2,说明defer 的函数是在return之前执行的
}
func f1() (r int64) {
	defer func() {
		r++
	}()
	r = 1
	return
}

//deferDemo3 defer 执行的时间
func deferDemo3() {
	fmt.Printf("r = %d\n", triple(5)) //r=15,说明defer 函数 在double函数之后执行
}

func double(x int64) int64 {
	return x + x
}

func triple(x int64) (r int64) {

	fmt.Printf("r------%d\n", r) //r------0
	defer func() {
		fmt.Printf("r+++++++++%d\n", r) //r+++++++++10

		r += x
		fmt.Printf("r======%d\n", r) //r======15
	}()
	fmt.Printf("r>>>>>>>>%d\n", r) //r>>>>>>>>0
	return double(x)
}
