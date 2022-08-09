package main

import (
	"fmt"
	"unsafe"
)

/**
*@author: 廖理
*@date:2022/8/9
**/

func structDemo() {
	d := new(Dog)

	size := unsafe.Sizeof(d) //d是一个引用
	fmt.Printf("unsafe.Sizeof(&d):%v\n", size)
	size = unsafe.Sizeof(*d)
	fmt.Printf("unsafe.Sizeof(d):%v\n", size)
	c := Cat{Animal{""}, 5}
	size = unsafe.Sizeof(&c)
	fmt.Printf("unsafe.Sizeof(&c):%v\n", size)
	size = unsafe.Sizeof(c)
	fmt.Printf("unsafe.Sizeof(c):%v\n", size)

	fmt.Printf("unsafe.Sizeof(int32):%v\n", unsafe.Sizeof(int32(12)))
	fmt.Printf("unsafe.Sizeof(int64):%v\n", unsafe.Sizeof(int64(12)))
	fmt.Printf("unsafe.Sizeof(d.age):%v\n", unsafe.Sizeof(d.age))
	fmt.Printf("unsafe.Sizeof(c.age):%v\n", unsafe.Sizeof(c.age))
	fmt.Printf("unsafe.Sizeof(string):%v\n", unsafe.Sizeof("faa"))
}
