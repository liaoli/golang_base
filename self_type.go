package main

import "fmt"

/**
*@author: 廖理
*@date:2022/11/16
**/
/**
*@author: 廖理
*@date:2022/11/16
**/

//Go语言基础之结构体
//发布于2017/06/23 ,更新于2017/06/23 22:08:00 | Golang |总阅读量：6771次
//Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
//
//类型别名和自定义类型
//自定义类型
//在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。
//
//自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：
//
////将MyInt定义为int类型
//type MyInt int
//通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。
//
//类型别名
//类型别名是Go1.9版本添加的新功能。
//
//类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。
//
//type TypeAlias = Type
//我们之前见过的rune和byte就是类型别名，他们的定义如下：
//
//type byte = uint8
//type rune = int32
//类型定义和类型别名的区别
//类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

//类型定义
type NewInt int

//类型别名
type MyInt = int

func selfTypeDemo() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}

//结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。
//MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。

func TypeDemo() {
	selfTypeDemo()
}
