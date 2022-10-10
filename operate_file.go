package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
*@author: 廖理
*@date:2022/8/9
**/

//结构体：
//是一种数据  类型。
//
//type Person struct {		—— 类型定义 （地位等价于 int byte bool string ....） 通常放在全局位置。
//	name string
//	sex  byte
//	age int
//}
//
//普通变量定义和初始化：
//
//1. 顺序初始化: 依次将结构体内部所欲成员初始化。
//
//var man Person = Person{"andy"， 'm',  20}
//
//2. 指定成员初始化：
//
//man := Person{name:"rose", age:18}		---- 未初始化的成员变量，取该数据类型对应的默认值
//
//普通变量的赋值和使用：
//
//使用“.”索引成员变量。
//
//var man3 Person
//man3.name = "mike"
//man3.sex = 'm'
//man3.age = 99
//
//结构体变量的比较：
//
//1. 比较： 只能使用 == 和 != 	不能 > < >= <=...
//
//2. 相同结构体类型（成员变量的类型、个数、顺序一致）变量之间可以直接赋值。
//
//结构体地址：
//
//结构体变量的地址 == 结构体首个元素的地址。
//
//结构体传参：
//unSafe.Sizeof(变量名) ——> 此种类型的变量所占用的内存空间大小
//
//将结构体变量的值拷贝一份，传递。	—— 几乎不用。 内存消耗大，效率低。
//
//指针变量定义和初始化：
//
//1. 顺序初始化: 依次将结构体内部所欲成员初始化。
//
//var man *Person = &Person{"andy"， 'm',  20}
//
//2. new(Person)
//
//p := new(Person)
//p.name = "name"
//p.age = 10
//
//
//指针索引成员变变量：
//
//使用“.”索引成员变量。
//
//var man3 Person
//man3.name = "mike"
//man3.sex = 'm'
//man3.age = 99
//结构体地址：
//
//结构体指针变量的值 == 结构体首个元素的地址。
//
//结构体指针传参：
//unSafe.Sizeof(指针) ： 不管何种类型的指针，在 64位操作系统下，大小一致。均为 8 字节！！！
//
//将结构体变量地址值，传递（传引用）。	—— 使用频率非常高！！！
//
//练习：
//定义一个结构体，包含成员 string、int、bool、[]string.
//
//在main函数中定义结构体 “普通变量”， 不初始化。 封装函数 initFunc， 在该函数内初始化 ， main 函数中打印查看。
//
//
//结构体指针做函数返回值：
//
//不能返回局部变量的地址。—— 局部变量保存栈帧上，函数调用结束后，栈帧释放。局部变量的地址，不再受系统保护，随时可能分配给其他程序。
//
//可以返回局部变量的值。
//
//字符串处理函数：
//
//1. 字符串按 指定分割符拆分：	Split
//
//ret := strings.Split(str, " I")
//
//2. 字符串按 空格拆分： Fields
//
//ret = strings.Fields(str)
//
//3. 判断字符串结束标记 HasSuffix
//
//flg := strings.HasSuffix("test.abc", ".mp3")
//
//4. 判断字符串起始标记 HasPrefix
//
//flg := strings.HasPrefix("test.abc", "tes.")
//
//打开、创建文件：
//
//1. 创建文件  Create： 	文件不存在创建， 文件存在，将文件内容清空。
//
//参数：name， 打开文件的路径： 绝对路径、相对路径		目录分割符：/
//
//2. 打开文件 Open：		以只读方式打开文件。文件不存在，打开失败。
//
//参数：name， 打开文件的路径： 绝对路径、相对路径
//
//3. 打开文件 OpenFile：	以只读、只写、读写 方式打开文件。文件不存在，打开失败。
//
//参1：name， 打开文件的路径： 绝对路径、相对路径
//
//参2：打开文件权限： O_RDONLY、O_WRONLY、O_RDWR
//
//参3：一般传 6
//
//写文件：
//
//按字符串写：WriteString（）	--> n个写入的字符个数
//
//n, err := f.WriteString("123")
//
//回车换行：	windows： \r\n	Linux: \n
//
//按位置写:
//Seek(): 	修改文件的读写指针位置。
//
//参1： 偏移量。 正：向文件尾偏， 负：向文件头偏
//
//参2： 偏移起始位置：
//
//io.SeekStart: 文件起始位置
//
//io.SeekCurrent： 文件当前位置
//
//io.SeekEnd: 文件结尾位置
//
//返回值：表示从文件起始位置，到当前文件读写指针位置的偏移量。
//
//off, _ := f.Seek(-5, io.SeekEnd)
//按字节写：
//writeAt():  在文件制定偏移位置，写入 []byte ,  通常搭配 Seek()
//
//参1： 待写入的数据
//
//参2：偏移量
//
//返回：实际写出的字节数。
//
//n, _ = f.WriteAt([]byte("1111"), off)
//
//读文件：
//按行读
//1）.  创建一个带有缓冲区的 Reader（读写器）
//
//reader : = bufio.NewReader(打开的文件指针)
//
//2）.  从reader的缓冲区中 ，读取指定长度的数据。数据长度取决于 参数 dlime
//
//buf, err := reader.ReadBytes( ' \n' )  	按行读。
//
//判断到达文件结尾： if err != nil && err == io.EOF  到文件结尾。
//
//文件结束标记，是要单独读一次获取到的。
//
//缓冲区：内存中的一块区域，用来减少物理磁盘访问操作。 《计算硬件及组成原理》 —— 机械工业出版社。
//
//按字节读、写文件。
//
//read([]byte):  按字节读文件
//
//write([]byte)：按字节字节
//
//目录操作：
//
//打开目录： OpenFile
//
//打开目录 OpenFile：	以只读方式打开目录。
//
//参1：name， 打开目录的路径： 绝对路径、相对路径
//
//参2：打开目录权限： O_RDONLY
//
//参3：os.ModeDir
//
//返回值： 返回一个可以读目录的 文件指针。
//
//读目录：Readdir
//
//原型：func (f *File) Readdir(n int) ([]FileInfo, error) {
//
//参数: 欲打开的目录项个数。 -1 ， 表所有
//
//	返回值：  FileInfo ：
//	type FileInfo interface {
//		Name() string       		// base name of the file
//		Size() int64        		// length in bytes for regular files; system-dependent for others
//		Mode() FileMode     		// file mode bits
//		ModTime() time.Time 	// modification time
//		IsDir() bool        		// abbreviation for Mode().IsDir()
//		Sys() interface{}   		// underlying data source (can return nil)
//	}
//
//
//
//	练习题3， 思路分析：
//
//	1. 根据用户指定的目录， 只读打开	—— 读目录的练习题。
//
//	2. 找到目录中的 .txt， 有可能有多个	—— 目录中找一个 指定类型文件
//
//	3. 打开 其中一个 .txt 文件。 循环读取一行。reader := bufio.NewReader, reader.ReadBytes('\n')   	—— 读一行文件内容练习题
//
//	4. 将一行数据的字符串，拆分后，存入 []string 。 Split、Fields	—— 字符串练习题
//
//	5. 遍历[]string 统计“Love”单词出现的次数。	—— map练习题
//
//C:/itcast/test

func FileDemo() {
	//dirDemo()
	//openFileDemo1()
	//openFileDemo()
	//readFile()

	copyFile()
}

func openFileDemo1() {
	f, err := os.Open("/Users/hfy/Desktop/goFileTest.txt") //只能打开已经存在的文件
	if err != nil {
		fmt.Println("打开文件错误 err:", err)
		return
	}
	defer f.Close()
	fmt.Println("打开文件成功 ")

	_, err = f.WriteString("go write file")

	if err != nil {
		fmt.Println("写文件错误 err:", err)
	}

}

func openFileDemo() { //
	f, err := os.OpenFile("/Users/hfy/Desktop/goFileTest.txt", os.O_RDWR|os.O_APPEND, 6) //os.O_APPEND在原来的内容后面追加 os.O_RDWR会直接覆盖之前的内容
	if err != nil {
		fmt.Println("打开文件错误 err:", err)
		return
	}
	defer f.Close()
	fmt.Println("打开文件成功 ")

	_, err = f.WriteString("go write file")

	if err != nil {
		fmt.Println("写文件错误 err:", err)
		return
	}
	fmt.Println("写文件成功")
}

func readFile() {

	f, err := os.OpenFile("/Users/hfy/Desktop/as_aftersale_order_multi_model.yaml", os.O_RDWR, 6)

	if err != nil {
		fmt.Println("打开文件错误 err:", err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		buf, err := reader.ReadBytes('\n')

		if err != nil && err == io.EOF {
			fmt.Println(string(buf))
			fmt.Println("读取文本完毕")
			return
		} else if err != nil {
			fmt.Println("读取文本完毕错误 err:", err)
		}
		fmt.Println(string(buf))
	}

}

func copyFile() {
	f_r, err := os.Open("/Users/hfy/Desktop/as_aftersale_order_multi_model.yaml")

	if err != nil {
		fmt.Println("打开文件错误 err:", err)
		return
	}
	defer f_r.Close()

	f_w, err := os.Create("/Users/hfy/Desktop/as_aftersale_order_multi_model_bak.yaml")

	if err != nil {
		fmt.Println("创建文件错误 err:", err)
		return
	}
	defer f_w.Close()

	buf := make([]byte, 2048)
	for {

		n, err := f_r.Read(buf)
		if err != nil {
			if err == io.EOF {
				f_w.Write(buf[:n])
				fmt.Printf("读取文件完毕,n=%d\n，复制文件完成", n)
				return
			}
			fmt.Printf("写入文件失败文件完毕,n=%d\n", n)

		}
		f_w.Write(buf[:n])
	}

}
func dirDemo() {

	f, err := os.OpenFile("/Users/hfy/Desktop", os.O_RDONLY, os.ModeDir)

	if err != nil {
		fmt.Println("打开文件夹错误 err:", err)
		return
	}
	defer f.Close()

	//读取目录
	info, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("读取文件夹错误 err:", err)
		return
	}

	for _, v := range info {
		if v.IsDir() {
			fmt.Println(v.Name(), "（文件夹）")
		} else {
			fmt.Println(v.Name(), "（文件）")
		}
	}

}
