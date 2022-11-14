package main

import "fmt"

/**
*@author: 廖理
*@date:2022/11/14
**/

//数组的初始化也有很多方式。
//
//方法一
//初始化数组时可以使用初始化列表来设置数组元素的值。

func arrayInit() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}

//方法二
//按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：

func arrayInit2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

//方法三
//我们还可以使用指定索引值的方式来初始化数组，例如:

func arrayInit3() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

//多维数组
//Go语言是支持多维数组的，我们这里以二维数组为例（数组中又嵌套数组）。
//
//二维数组的定义
func multidimensionalArrayInit() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆

	//支持的写法
	c := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	//d := [3][...]string{
	//	{"北京", "上海"},
	//	{"广州", "深圳"},
	//	{"成都", "重庆"},
	//}

	fmt.Println(c)
	//fmt.Println(d)
}

//二维数组的遍历
func iterateMultidimensionalArray() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

//数组是值类型
//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

func modifyArray1(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func modifyArray() {
	a := [3]int{10, 20, 30}
	modifyArray1(a) //在modify中修改的是a的副本x
	fmt.Println(a)  //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

func ArrayDemo() {
	//arrayInit()
	//arrayInit2()
	//arrayInit3()
	//multidimensionalArrayInit()
	//iterateMultidimensionalArray()
	modifyArray()
}
