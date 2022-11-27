package main

import (
	"fmt"
	"unsafe"
)

/**
*@author: 廖理
*@date:2022/11/16
**/
//结构体
//Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，
//这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，
//可以封装多个基本数据类型，这种数据类型叫结构体，英文名称struct。 也就是我们可以通过struct来定义自己的类型了。
//
//Go语言中通过struct来实现面向对象。
//
//结构体的定义
//使用type和struct关键字来定义结构体，具体代码格式如下：
//
//type 类型名 struct {
//	字段名 字段类型
//	字段名 字段类型
//	…
//}
//其中：
//
//类型名：标识自定义结构体的名称，在同一个包内不能重复。
//字段名：表示结构体字段名。结构体中的字段名必须唯一。
//字段类型：表示结构体字段的具体类型。
//举个例子，我们定义一个Person（人）结构体，代码如下：
//
//type person struct {
//	name string
//	city string
//	age  int8
//}
//同样类型的字段也可以写在一行，
//
//type person1 struct {
//	name, city string
//	age        int8
//}
//这样我们就拥有了一个person的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄。
//这样我们使用这个person结构体就能够很方便的在程序中表示和存储人信息了。
//
//语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

//结构体实例化
//只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
//
//结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
//
//var 结构体实例 结构体类型
//基本实例化
//举个例子：

type person struct {
	name string
	city string
	age  int8
}

func StructInit() {
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
}

//我们通过.来访问结构体的字段（成员变量）,例如p1.name和p1.age等。

//匿名结构体
//在定义一些临时数据结构等场景下还可以使用匿名结构体。
//

func NONameStruct() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

//创建指针类型结构体
//我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：
//
//var p2 = new(person)
//fmt.Printf("%T\n", p2)     //*main.person
//fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
//从打印的结果中我们可以看出p2是一个结构体指针。
//
//需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。

//创建指针类型结构体
//我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：
func pointerDemo1() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
}

//从打印的结果中我们可以看出p2是一个结构体指针。

//需要注意的是在Go语言中支持对结构体指针直接使用.来访问结构体的成员。
func pointerDemo2() {
	var p2 = new(person)
	p2.name = "小王子"
	p2.age = 28
	p2.city = "上海"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"小王子", city:"上海", age:28}
}

//取结构体的地址实例化
//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
func pointerDemo3() {
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}
}

//结构体初始化
//没有初始化的结构体，其成员变量都是对应其类型的零值。

//使用键值对初始化
//使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

func initStruct() {
	//使用键值对初始化
	//使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。
	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}
	//也可以对结构体指针进行键值对初始化，例如：

	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}
	//当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

	p7 := &person{
		city: "北京",
	}
	fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
	//使用值的列表初始化
	//初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：

	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}
	//使用这种格式初始化时，需要注意：
	//
	//必须初始化结构体的所有字段。
	//初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	//该方式不能和键值初始化方式混用。

}

//结构体内存布局
//结构体占用一块连续的内存。

type testMemory struct {
	a int8
	b int8
	c int8
	d int8
}

func StructMemory() {
	n := testMemory{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)
	fmt.Printf("n.b %p\n", &n.b)
	fmt.Printf("n.c %p\n", &n.c)
	fmt.Printf("n.d %p\n", &n.d)
	//输出：

	//n.a 0xc0000200a0
	//n.b 0xc0000200a1
	//n.c 0xc0000200a2
	//n.d 0xc0000200a3

	//【进阶知识点】关于Go语言中的内存对齐推荐阅读:Go结构体的内存对齐
}

//空结构体
//空结构体是不占用空间的。

var v struct{}

func EmptyStructMemory() {
	fmt.Println(unsafe.Sizeof(v)) // 0
}

//面试题
//请问下面代码的执行结果是什么？

type student struct {
	name string
	age  int
}

func interview() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Printf("k%s=>v%#p\n", k, v)
	}
}

//构造函数
//Go语言的结构体没有构造函数，我们可以自己实现。 例如，下方的代码就实现了一个person的构造函数。
//因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

//调用构造函数
func CreateFuc() {
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}
}

//方法和接收者
//Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。
//接收者的概念就类似于其他语言中的this或者 self。
//
//方法的定义格式如下：
//
//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	函数体
//}
//其中，
//
//接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。
//例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
//接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
//方法名、参数列表、返回参数：具体格式与函数定义相同。
//举个例子：

//Dream Person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func StructFunc() {
	p1 := newPerson("小王子", "深圳", 25)
	p1.Dream()
}

//方法与函数的区别是，函数不属于任何类型，方法属于特定的类型。
//
//指针类型的接收者
//指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。
//这种方式就十分接近于其他语言中面向对象中的this或者self。 例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。

// SetAge 设置p的年龄
// 使用指针接收者
func (p *person) SetAge(newAge int8) {
	p.age = newAge
}

//调用该方法：

func StructFuncPoint() {
	p1 := newPerson("小王子", "", 25)
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30
}

//值类型的接收者
//当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

// SetAge2 设置p的年龄
// 使用值接收者
func (p person) SetAge2(newAge int8) {
	p.age = newAge
}

func StructFunc2() {
	p1 := newPerson("小王子", "", 25)
	p1.Dream()
	fmt.Println(p1.age) // 25
	p1.SetAge2(30)      // (*p1).SetAge2(30)
	fmt.Println(p1.age) // 25
}

//什么时候应该使用指针类型接收者
//需要修改接收者中的值
//接收者是拷贝代价比较大的大对象
//保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

func StructDemo() {
	//StructInit()

	//NONameStruct()
	//pointerDemo1()
	//pointerDemo2()
	//pointerDemo3()
	//initStruct()
	//StructMemory()
	//EmptyStructMemory()
	//interview()
	//CreateFuc()
	StructFunc()
}
