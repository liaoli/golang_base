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

//从上面的结果可以看到我们的测试用例覆盖了100%的代码。
//
//Go还提供了一个额外的-coverprofile参数，用来将覆盖率相关的记录信息输出到一个文件。例如：

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±  go test -cover -coverprofile=c.out
//PASS
//coverage: 100.0% of statements
//ok      awesomeProject/unit_test        0.339s
//上面的命令会将覆盖率相关的信息输出到当前文件夹下面的c.out文件中，
//然后我们执行go tool cover -html=c.out，使用cover工具来处理生成的记录信息，
//该命令会打开本地的浏览器窗口生成一个HTML报告。

//基准测试函数格式
//基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
//
//func BenchmarkName(b *testing.B){
//	// ...
//}
//基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，
//这样的测试才有对照性，b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。
//go test -bench=Split
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unit_test.Split("沙河有沙又有河", "沙")
	}
}

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±  go test -bench=Split
//goos: darwin
//goarch: amd64
//pkg: awesomeProject/unit_test
//cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
//BenchmarkSplit-12       10289245               114.7 ns/op
//PASS
//ok      awesomeProject/unit_test        2.649s

//其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
//10289245和114.7ns/op表示每次调用Split函数耗时114.7ns，这个结果是10289245次调用的平均值。
//
//我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±  go test -bench=Split -benchmem
//goos: darwin
//goarch: amd64
//pkg: awesomeProject/unit_test
//cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
//BenchmarkSplit-12       10300904               114.6 ns/op            48 B/op          2 allocs/op
//PASS
//ok      awesomeProject/unit_test        1.628s

//其中，48 B/op表示每次操作内存分配了48字节，2 allocs/op则表示每次操作进行了2次内存分配。

func BenchmarkSplit2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unit_test.Split2("沙河有沙又有河", "沙")
	}
}

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±  go test -bench=Split2 -benchmem
//goos: darwin
//goarch: amd64
//pkg: awesomeProject/unit_test
//cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
//BenchmarkSplit2-12      10246958               115.0 ns/op            48 B/op          1 allocs/op
//PASS
//ok      awesomeProject/unit_test        1.636s
