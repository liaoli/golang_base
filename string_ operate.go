package main

import (
	"fmt"
	"strings"
)

/**
*@author: 廖理
*@date:2022/8/5
**/

//在对字符串进行处理时，需要借助于包“strings”
//次包里面提供了大量的操作字符串的函数，以后要处理字符串可以借助次工具包
//（1）Contains

func stringDemo() {

}

func fieldDemo() {
	list := strings.Fields("I love you") //根据空格划分生成切片
	fmt.Printf("Fields: %v\n", list)     //Love You More
}

func titleDemo() {
	title := strings.Title("love you more")
	fmt.Printf("Title: %s\n", title) //Love You More
}

func countDemo() {
	count := strings.Count("abcab", "ab")
	fmt.Printf("Count: %d\n", count)
}

func joinDemo() {
	result := strings.Join([]string{"a", "b", "c"}, ",")

	fmt.Printf("join: %s\n", result)
}

func containsDemo() {
	//字符串是否包含某个字符串
	var b1 bool = strings.Contains("abc", "abd")
	fmt.Printf("Contains:%t\n", b1)
}

func indexDemo() {
	index := strings.Index("abca", "a")

	fmt.Printf("Index: %d\n", index) //0

	index = strings.Index("abc", "g")

	fmt.Printf("Index: %d\n", index) //-1

	lastIndex := strings.LastIndex("abca", "a") //3

	fmt.Printf("LastIndex: %d\n", lastIndex)

	lastIndexAny := strings.LastIndexAny("abca", "a") //3

	fmt.Printf("LastIndexAny: %d\n", lastIndexAny)
}
