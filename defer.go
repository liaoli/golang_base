package main

import "fmt"

/**
*@author: 廖理
*@date:2022/11/15
**/

//defer语句
//Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，
//将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
//
//举个例子：

func deferTest1() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

//输出结果：
//start
//end
//3
//2
//1

//由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

//defer执行时机
//在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
//而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。具体如下图所示：

//defer经典案例
//阅读下面的代码，写出最后的打印结果。

func fc1() int {
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x //1。返回值赋值，2.defer 3. return RET指令
}

func fc2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值 x= 5 2.defer x=6 3.RET x = 6
}

func fc3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //1.y = x = 5 2.defer x = 5 3.RET  y = 5
}
func fc4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 //1.x =5 //2defer x值传递，返回值的x 不会受影响 3.RET x = 5
}

func fc5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5 //1.x =5 //2defer x值传递，返回值的x 不会受影响 3.RET x = 5
}

func deferTest2() {
	fmt.Println(fc1())
	fmt.Println(fc2())
	fmt.Println(fc3())
	fmt.Println(fc4())
	fmt.Println(fc5())
}

//defer面试题
func calcx(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//
func deferTest3() {
	x := 1
	y := 2
	defer calcx("AA", x, calcx("A", x, y))
	x = 10
	defer calcx("BB", x, calcx("B", x, y))
	y = 20
}

//1.值传递和引用传递的区别

//
//问，上面代码的输出结果是？（提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值）
func deferTestDemo() {
	//deferTest1()
	//deferTest2()
	//deferTest3()
}

type T int

func (t T) M(n int) T {
	fmt.Println(n)
	return t
}

func DeferDemo() {
	var t T
	// t.M(1)是方法调用M(2)的属主实参，因此它
	// 将在M(2)调用被推入延迟调用队列时被估值。
	defer t.M(1).M(2)
	t.M(3).M(4)
}
