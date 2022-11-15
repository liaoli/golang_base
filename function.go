package main

import (
	"errors"
	"fmt"
	"strings"
)

/**
*@author: 廖理
*@date:2022/11/15
**/
//
//函数是组织好的、可重复使用的、用于执行指定任务的代码块。本文介绍了Go语言中函数的相关内容。
//
//函数
//Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”。
//
//函数定义
//Go语言中定义函数使用func关键字，具体格式如下：
//
//func 函数名(参数)(返回值){
//	函数体
//}
//其中：
//
//函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名（包的概念详见后文）。
//参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
//返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用()包裹，并用,分隔。
//函数体：实现指定功能的代码块。
//我们先来定义一个求两个数之和的函数：

func intSum(x int, y int) int {
	return x + y
}

//函数的参数和返回值都是可选的，例如我们可以实现一个既不需要参数也没有返回值的函数：

func sayHello() {
	fmt.Println("Hello 沙河")
}

//参数
//类型简写
//函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：
//
//func intSum(x, y int) int {
//	return x + y
//}
//上面的代码中，intSum函数有两个参数，这两个参数的类型均为int，因此可以省略x的类型，因为y后面有类型说明，x参数也是该类型。
//
//可变参数
//可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
//
//注意：可变参数通常要作为函数的最后一个参数。
//
//举个例子：

func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例代码如下：

func intSum3(x int, y ...int) int {
	fmt.Println(x, y)

	fmt.Printf("y 的类型%T\n", y) //本质上，函数的可变参数是通过切片来实现的。
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//返回值
//Go语言中通过return关键字向外输出返回值。
//
//多返回值
//Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
//
//举个例子：

func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名
//函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
//
//例如：

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//返回值补充
//当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

func someFunc(x string) []int {
	if x == "" {
		return nil // 没必要返回[]int{}
	}
	//...

	return make([]int, 0)
}

//函数进阶
//变量作用域
//全局变量
//全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}

//func main() {
//	testGlobalVar() //num=10
//}
//局部变量
//局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：

func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

//func main() {
//	testLocalVar()
//	fmt.Println(x) // 此时无法使用变量x
//}
//如果局部变量和全局变量重名，优先访问局部变量。

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}

//接下来我们来看一下语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。

func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}

//还有我们之前讲过的for循环语句中定义的变量，也是只在for语句块中生效：

func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}

//函数类型与变量
//定义函数类型
//我们可以使用type关键字来定义一个函数类型，具体格式如下：
//
//type calculation func(int, int) int
//上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。
//
//简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。

type calculation func(int, int) int

func addCal(x, y int) int {
	return x + y
}

func subCal(x, y int) int {
	return x - y
}

//add和sub都能赋值给calculation类型的变量。
//
//var c calculation
//c = add
//函数类型变量
//我们可以声明函数类型的变量并且为该变量赋值：

func FunctionVariable() {
	var c calculation               // 声明一个calculation类型的变量c
	c = addCal                      // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := subCal                     // 将函数add赋值给变量f
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}

//高阶函数
//高阶函数分为函数作为参数和函数作为返回值两部分。
//
//函数作为参数
//函数可以作为参数：

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func functionParam() {
	ret2 := calc(10, 20, addCal)
	fmt.Println(ret2) //30
}

//函数作为返回值
//函数也可以作为返回值：

func functionReturn(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return addCal, nil
	case "-":
		return subCal, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func functionReturnDemo() {
	add, err := functionReturn("+")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(add(1, 2))
	}
}

//匿名函数和闭包
//匿名函数
//函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：
//
//func(参数)(返回值){
//	函数体
//}
//匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

func CloserDemo1() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

//匿名函数多用于实现回调函数和闭包。

//闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 首先我们来看一个例子：

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func CloserDemo2() {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}

//变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。 在f的生命周期内，变量x也一直有效。 闭包进阶示例1：

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func CloserDemo3() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}

//闭包进阶示例2：

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func CloserDemo4() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

//闭包进阶示例3：

func CloserCalc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func CloserDemo5() {
	f1, f2 := CloserCalc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

//闭包其实并不复杂，只要牢记闭包=函数+引用环境。

func FunctionDemo() {
	//intSum3(1, 2, 3, 4, 5)
	//testNum()
	//FunctionVariable()
	//functionParam()
	//functionReturnDemo()
	//CloserDemo1()
	//CloserDemo2()
	//CloserDemo3()
	//CloserDemo4()
	//CloserDemo5()
	panicDemo()
}

//内置函数介绍
//内置函数	介绍
//close	主要用来关闭channel
//len	用来求长度，比如string、array、slice、map、channel
//new	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
//make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
//append	用来追加元素到数组、slice中
//panic和recover	用来做错误处理
//panic/recover
//Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。 首先来看一个例子：

//内置函数介绍
//内置函数	介绍
//close	主要用来关闭channel
//len	用来求长度，比如string、array、slice、map、channel
//new	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
//make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
//append	用来追加元素到数组、slice中
//panic和recover	用来做错误处理
//panic/recover
//Go语言中目前（Go1.12）是没有异常机制，但是使用panic/recover模式来处理错误。 panic可以在任何地方引发，但recover只有在defer调用的函数中有效。 首先来看一个例子：

func funcA1() {
	fmt.Println("func A")
}

func funcB1() {
	panic("panic in B")
}

func funcC1() {
	fmt.Println("func C")
}
func panicDemo() {
	funcA1()
	funcB1()
	funcC1()
}
