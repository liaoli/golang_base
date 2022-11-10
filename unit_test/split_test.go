package unit_test_test

import (
	"awesomeProject/unit_test"
	"reflect"
	"testing"
)

/**
*@author: 廖理
*@date:2022/11/10
**/

func TestSplit(t *testing.T) {
	got := unit_test.Split("a:b:c", ":") // 程序输出的结果
	want := []string{"a", "b", "c"}      // 期望的结果
	if !reflect.DeepEqual(want, got) {   // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

//一个测试用例有点单薄，我们再编写一个测试使用多个字符切割字符串的例子，在split_test.go中添加如下测试函数：
func TestSplitWithComplexSep(t *testing.T) {
	got := unit_test.Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSplitWithGroup(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := unit_test.Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%#v, got:%#v", tc.want, got)
		}
	}
}

//我们都知道可以通过-run=RegExp来指定运行的测试用例，
//还可以通过/来指定要运行的子测试用例，
//例如：go test -v -run=Split/simple只会运行simple对应的子测试用例。
func TestSplitWithSub(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := unit_test.Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

//测试覆盖率是你的代码被测试套件覆盖的百分比。
//通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。
//Go提供内置功能来检查你的代码覆盖率。我们可以使用go test -cover来查看测试覆盖率。
//
//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±  go test -cover
//PASS
//coverage: 100.0% of statements
//ok      awesomeProject/unit_test        0.276s
