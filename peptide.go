package main

import "fmt"

/**
*@author: 廖理
*@date:2022/8/5
**/

type Human interface {
	sayHi()
}

type Man struct {
	name string
}

func (m *Man) sayHi() {
	fmt.Printf("hi! I am a man..\n")
}

type Woman struct {
}

func (w *Woman) sayHi() {
	fmt.Printf("hi! I am a woman..\n")
}

//go 的多肽实现1
func testPeptide() {

	var m = Man{"li"}
	humanSayHi(&m)

	w := Woman{}

	humanSayHi(&w)

	var h Human
	//h =w 错误
	h = &w //正确
	h.sayHi()

}

func humanSayHi(h Human) {
	h.sayHi()
}

//go 的多肽实现2 错误 不能实现
func testPeptide2() {
	//animal := make([]*Animal,0)
	//d := Dog{Animal{"dog"}}
	//c := Cat{Animal{"dog"}}
	//
	//animal = append(animal,&d)//报错
	//animal = append(animal,d)//报错
}

type Animal struct {
	category string
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}
