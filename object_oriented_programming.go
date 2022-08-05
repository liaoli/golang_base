package main

import "fmt"

/**
*@author: 廖理
*@date:2022/8/5
**/

//当然严格意义上，GO语言中是没有继承的，
//但是我们可以通过”匿名组合”来实现继承的效果。

type Person struct {
	id   int64
	name string
	age  int64
}

type Student struct {
	//person Person//不是匿名的不可以直接引用
	Person //必须匿名才能直接引用
	score  int64
	name   string
}

func oopDemo() {
	var st1 Student
	fmt.Printf("结构体默认值：%v\n", st1)
	st1.age = 10
	st1.name = "tom"
	fmt.Printf("结构体复制后：%v\n", st1)
	var st Student = Student{Person{1, "lily", 25}, 90, "cc"}
	fmt.Printf("结构体初始化赋值：%v\n", st)

	//通过结果发现是对Student中的name进行赋值，所以在操作同名字段时，
	//有一个基本的原则：如果能够在自己对象所属的类（结构体）中找到对应的成员，
	//那么直接进行操作，如果找不到就去对应的父类（结构体）中查找。这就是所谓的就近原则。
	st.name = "同名字段" //

	fmt.Printf("结构体通明字段 st.name：%v\n", st.name)
	fmt.Printf("结构体通明字段 st.Person.name：%v\n", st.Person.name)

}

//指针匿名字段
func referencePointer() {

	var t1 teacher = teacher{}

	fmt.Printf("指针匿名字段结构体默认值 t1 = %v \n", t1)

	//t1.age = 100 //不能直接使用应为 此时的匿名应用字段的值是空 会报异常：panic: runtime error: invalid memory address or nil pointer dereference
	//必须先给匿名引用赋值之后才能操作

	t1.Person = &Person{100, "", 25} //想给匿名引用赋值才能直接使用

	t1.age = 100

	fmt.Printf("指针匿名字段结构体赋值 t1 = %v \n", t1)

	t2 := teacher{&Person{22, "kate", 22}, "chinese"}

	fmt.Printf("指针匿名字段结构体 t2 = %v \n", t2)

	//t1.test1()
	//t2.test2()

	t1.say() //同名方法和同名字段的用法是一样的就近原则
}

type teacher struct {
	*Person
	subject string
}

func (p Person) say() {
	fmt.Printf("我是一个人，我的名字叫%s,我今年%d\n", p.name, p.age)
}

//结构体方法 接收的是实体
func (t teacher) test1() {
	fmt.Printf("我是一名老师，我的名字叫%s,我教的科目是%s\n", t.name, t.subject)
}

//结构体方法 接收的是实体
func (t teacher) say() {
	fmt.Printf("我是一名老师，我的名字叫%s,我教的科目是%s\n", t.name, t.subject)
}

//结构体方法 接收的是引用
func (t *teacher) test2() {
	fmt.Printf("我是一名老师，我的名字叫%s,我教的科目是%s\n", t.name, t.subject)
}
