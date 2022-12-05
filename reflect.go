package main

import (
	"fmt"
	"reflect"
)

/**
*@author: 廖理
*@date:2022/12/5
**/

//Go语言中的变量是分为两部分的:
//
//类型信息：预先定义好的元信息。
//值信息：程序运行过程中可动态变化的。
//反射介绍
//反射是指在程序运行期间对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
//
//支持反射的语言可以在程序编译期间将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改它们。
//
//Go程序在运行期间使用reflect包访问程序的反射信息。
//
//在上一篇博客中我们介绍了空接口。 空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？ 反射就是在运行时动态的获取一个变量的类型信息和值信息。
//
//reflect包
//在Go语言的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的(我们在上一篇接口的博客中有介绍相关概念)。 在Go语言中反射的相关功能由内置的reflect包提供，任意接口值在反射中都可以理解为由reflect.Type和reflect.Value两部分组成，并且reflect包提供了reflect.TypeOf和reflect.ValueOf两个函数来获取任意对象的Value和Type。
//
//TypeOf
//在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func reflectDemo() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}

//type name和type kind
//在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。因为在Go语言中我们可以使用type关键字构造很多自定义类型，
//而种类（Kind）就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。
//举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。

type myInt int64

func reflectType1(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func reflectDemo1() {
	var a *float32  // 指针
	var b myInt     // 自定义类型
	var c rune      // 类型别名
	reflectType1(a) // type: kind:ptr
	reflectType1(b) // type:myInt kind:int64
	reflectType1(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var d = person{
		name: "沙河小王子",
		age:  18,
	}
	var e = book{title: "《跟小王子学Go语言》"}
	reflectType1(d) // type:person kind:struct
	reflectType1(e) // type:book kind:struct
}

//ValueOf
//reflect.ValueOf()返回的是reflect.Value类型，其中包含了原始值的值信息。reflect.Value与原始值之间可以互相转换。
//
//reflect.Value类型提供的获取原始值的方法如下：
//
//方法	说明
//Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
//Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
//Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
//Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
//Bool() bool	将值以 bool 类型返回
//Bytes() []bytes	将值以字节数组 []bytes 类型返回
//String() string	将值以字符串类型返回
//通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func reflectValueDemo() {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100
	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}

//通过反射设置变量的值
//想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
func reflectSetValuedemo() {
	var a int64 = 100
	//reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
