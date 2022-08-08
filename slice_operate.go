package main

import "fmt"

/**
*@author: 廖理
*@date:2022/8/5
**/

//
//指针：
//
//指针就是地址。 指针变量就是存储地址的变量。
//
//*p ： 解引用、间接引用。
//
//栈帧：	用来给函数运行提供内存空间。 取内存于 stack 上。
//
//当函数调用时，产生栈帧。函数调用结束，释放栈帧。
//
//栈帧存储： 1. 局部变量。 2. 形参。 （形参与局部变量存储地位等同） 3. 内存字段描述值
//
//指针使用注意：
//
//空指针：未被初始化的指针。	var p *int		*p --> err
//
//野指针：被一片无效的地址空间初始化。
//
//格式化输出：
//
//%q： 以Go语言格式显示字符串。 默认带有“”符
//
//%v： 显示对应数据详细信息
//
//变量存储：
//
//等号 左边的变量，代表 变量所指向的内存空间。	（写）
//
//等号 右边的变量，代表 变量内存空间存储的数据值。	（读）
//
//指针的函数传参（传引用）。
//
//传地址（引用）：将形参的地址值作为函数参数传递。
//
//传值（数据据）：将实参的 值 拷贝一份给形参。
//
//传引用：	在A栈帧内部，修改B栈帧中的变量值。
//
//
//切片：
//
//为什么用切片：
//
//1. 数组的容量固定，不能自动拓展。
//
//2. 值传递。 数组作为函数参数时，将整个数组值拷贝一份给形参。
//
//在Go语言当，我们几乎可以在所有的场景中，使用 切片替换数组使用。
//
//切片的本质：
//
//不是一个数组的指针，是一种数据结构体，用来操作数组内部元素。	runtime/slice.go	type slice struct {
//	*p
//	len
//	切片的使用：										cap
//}
//数组和切片定义区别：
//
//创建数组时 [ ] 指定数组长度。
//
//创建切片时， [] 为空，或者 ...
//
//切片名称 [ low : high : max ]
//
//low: 起始下标位置
//
//high：结束下标位置	len = high - low
//
//容量：cap = max - low
//
//截取数组，初始化 切片时，没有指定切片容量时， 切片容量跟随原数组（切片）。
//
//s[:high:max] : 从 0 开始，到 high结束。（不包含）
//
//s[low:] :	从low 开始，到 末尾
//
//s[: high]:	从 0 开始，到 high结束。容量跟随原先容量。【常用】
//
//切片创建：
//
//1. 自动推导类型创建 切片。slice := []int {1, 2, 4, 6}
//
//2. slice := make([]int, 长度，容量)
//
//3. slice := make([]int, 长度）		创建切片时，没有指定容量， 容量== 长度。【常用】
//
//切片做函数参数 —— 传引用。（传地址）
//
//append：在切片末尾追加元素
//
//append(切片对象， 待追加元素）
//
//向切片增加元素时，切片的容量会自动增长。1024 以下时，一两倍方式增长。
//
//copy：
//
//copy（目标位置切片， 源切片）
//
//拷贝过程中，直接对应位置拷贝。
//
//map：
//
//字典、映射  	key —— value	key： 唯一、无序。 不能是引用类型数据。
//
//map 不能使用 cap（）
//
//创建方式：
//1.  var m1 map[int]string		--- 不能存储数据-
//
//2. m2 := map[int]string{}		---能存储数据
//
//3. m3 := make(map[int]string)		---默认len = 0
//
//4. m4 := make(map[int]string, 10)
//
//初始化：
//
//1. var m map[int]string = map[int]string{ 1: "aaa", 2:"bbb"}	保证key彼此不重复。
//
//2. m := map[int]string{ 1: "aaa", 2:"bbb"}
//
//赋值:
//
//赋值过程中，如果新map元素的key与原map元素key 相同 	——> 覆盖（替换）
//
//赋值过程中，如果新map元素的key与原map元素key 不同	——> 添加
//
//map的使用：
//
//遍历map：
//
//for  key值， value值 := range map {
//
//}
//
//for  key值 := range map {
//
//}
//
//判断map中key是否存在。
//
//map[下标] 运算：返回两个值， 第一个表 value 的值，如果value不存在。 nil
//
//第二个表 key是否存在的bool类型。存在 true， 不存在false
//
//删除map：
//
//delete()函数： 	参1： 待删除元素的map	参2： key值
//
//delete（map， key）	删除一个不存在的key ， 不会报错。
//
//map 做函数参数和返回值，传引用。

//要删除slice中间的某个元素并保存原有的元素顺序， 如：
//{5, 6, 7, 8, 9} ——> {5, 6, 8, 9}

func deleteSliceItem() {
	slice := []int64{5, 6, 7, 8, 9}
	fmt.Printf("删除前：%v\n", slice)

	dest := slice[2:]

	src := slice[3:]
	fmt.Printf("dest：%v\n", dest)
	fmt.Printf("src：%v\n", src)

	copy(dest, src)

	result := slice[:len(slice)-1]

	fmt.Printf("删除后dest：%v\n", dest)
	fmt.Printf("删除后slice：%v\n", slice)
	fmt.Printf("删除后result：%v\n", result)

}

func remove(data []int, idx int) []int {
	copy(data[idx:], data[idx+1:])
	return data[:len(data)-1]
}
