package main

import "fmt"

/**
*@author: 廖理
*@date:2022/8/5
**/

//空接口(interface{})不包含任何的方法，正因为如此，
//所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。上帝

type A struct {
}

type B struct {
}

func testEmptyInterface() {
	data := make([]interface{}, 0)

	data = append(data, A{})
	data = append(data, B{})

	for _, v := range data {
		fmt.Printf("v = %T\n", v)
	}
}
