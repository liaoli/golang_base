package main

import "fmt"

/**
*@author: 廖理
*@date:2022/8/5
**/

//切片
//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
//
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
//
//切片的定义
//声明切片类型的基本语法如下：
//
//var name []T
//其中，
//
//name:表示变量名
//T:表示切片中的元素类型
//举个例子：

func initSlice() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	//var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	//fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

//切片的长度和容量
//切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
//
//切片表达式
//切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种指定low和high两个索引界限值的简单的形式，
//另一种是除了low和high索引界限值外还指定容量的完整的形式。
//
//简单切片表达式
//切片的底层就是一个数组，所以我们可以基于数组通过切片表达式得到切片。 切片表达式中的low和high表示一个索引范围（左包含，右不包含），
//也就是下面代码中从数组a中选出1<=索引值<4的元素组成切片s，得到的切片长度=high-low，容量等于得到的切片的底层数组的容量。

func sliceLenthAndCap() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}

//为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
//
//a[2:]  // 等同于 a[2:len(a)]
//a[:3]  // 等同于 a[0:3]
//a[:]   // 等同于 a[0:len(a)]
//注意：
//
//对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
//
//对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。常量索引必须是非负的，并且可以用int类型的值表示;
//对于数组或常量字符串，常量索引也必须在有效范围内。如果low和high两个指标都是常数，它们必须满足low <= high。如果索引在运行时超出范围，
//就会发生运行时panic。

func SliceSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))
}

//完整切片表达式
//对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：
//
//a[low : high : max]
//上面的代码会构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片。
//另外，它会将得到的结果切片的容量设置为max-low。在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0。

func AllSlice() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[1:3:4]
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))
}

//使用make()函数构造切片
//我们上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的make()函数，格式如下：
//
//make([]T, size, cap)
//其中：
//
//T:切片的元素类型
//size:切片中元素的数量
//cap:切片的容量
//举个例子：

func makeSlice() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}

//上面代码中a的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以len(a)返回2，cap(a)则返回该切片的容量。

//切片的本质
//切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

//判断切片是否为空
//要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
//
//切片不能直接比较
//切片之间是不能比较的，我们不能使用==操作符来判断两个切片是否含有全部相等元素。 切片唯一合法的比较操作是和nil比较。
//一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是nil，例如下面的示例：

//var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
//s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
//s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil
//所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断。

//切片的赋值拷贝
//下面的代码中演示了拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

func initCopy() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}

//使用copy()函数复制切片

//由于切片是引用类型，所以a和b其实都指向了同一块内存地址。修改b的同时a的值也会发生变化。
//
//Go语言内建的copy()函数可以迅速地将一个切片的数据复制到另外一个切片空间中，copy()函数的使用格式如下：
//
//copy(destSlice, srcSlice []T)
//其中：
//
//srcSlice: 数据来源切片
//destSlice: 目标切片
//举个例子：

func SliceCopy() {
	// copy()复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 5)
	copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1 2 3 4 5]
	c[0] = 1000
	fmt.Println(a) //[1 2 3 4 5]
	fmt.Println(c) //[1000 2 3 4 5]
}

//从切片中删除元素
//Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

func deleteSlice() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}

//总结一下就是：要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
func SliceDemo() {
	//initSlice()
	//sliceLenthAndCap()
	//SliceSlice()
	//AllSlice()
	//makeSlice()
	//initCopy()
	//SliceCopy()
	deleteSlice()
}

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
