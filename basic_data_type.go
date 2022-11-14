package main

import (
	"fmt"
	"unsafe"
)

/**
*@author: 廖理
*@date:2022/11/14
**/

func intType() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF
	fmt.Printf("%d \n", c) // FF
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
		//	104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 178(²) 153 230(æ) 178(²) 179(³)
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r) //104(h) 101(e) 108(l) 108(l) 111(o) 27801(沙) 27827(河)
	}
	fmt.Println()
}

func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

//编写代码统计出字符串"hello沙河小王子"中汉字的数量。
//4E00-9FA5
func test() {
	s := "hello沙c334河1小is王not 子 a "
	var hz int64
	hzList := make([]rune, 0)
	for _, v := range s {

		if v > 0x4E00 && v < 0x9FA5 {
			hz++
			hzList = append(hzList, v)
		}
	}
	fmt.Printf("汉字个数为%d \n", hz)
	fmt.Printf("rune的字节数为%d \n", unsafe.Sizeof(hzList[1]))
	fmt.Printf("汉字为:%s \n", string(hzList))
}

func DataType() {
	//traversalString()
	//changeString()
	test()
}
