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
