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

//性能比较函数
//上面的基准测试只能得到给定操作的绝对耗时，但是在很多性能问题是发生在两个不同操作之间的相对耗时，
//比如同一个函数处理1000个元素的耗时与处理1万甚至100万个元素的耗时的差别是多少？再或者对于同一个
//任务究竟使用哪种算法性能最佳？我们通常需要对两个不同算法的实现使用相同的输入来进行基准比较测试。
//
//性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。
//

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		unit_test.Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±✚   go test -bench=.
//goos: darwin
//goarch: amd64
//pkg: awesomeProject/unit_test
//cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
//BenchmarkFib1-12        737591203                1.628 ns/op
//BenchmarkFib2-12        263815970                4.551 ns/op
//BenchmarkFib3-12        160421407                7.477 ns/op
//BenchmarkFib10-12        4211372               285.0 ns/op
//BenchmarkFib20-12          32959             36359 ns/op
//BenchmarkFib40-12              2         536245811 ns/op
//BenchmarkSplit-12       10298650               114.0 ns/op
//BenchmarkSplit2-12      10613989               112.1 ns/op
//PASS
//ok      awesomeProject/unit_test        12.633s

//这里需要注意的是，默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，
//则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
//
//最终的BenchmarkFib40只运行了两次，每次运行的平均值只有不到一秒。
//像这种情况下我们应该可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。
//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±✚  go test -bench=Fib40 -benchtime=20s
//goos: darwin
//goarch: amd64
//pkg: awesomeProject/unit_test
//cpu: Intel(R) Core(TM) i5-10400 CPU @ 2.90GHz
//BenchmarkFib40-12             44         536031844 ns/op
//PASS
//ok      awesomeProject/unit_test        24.580s

//这一次BenchmarkFib40函数运行了44次，结果就会更准确一些了。
//
//使用性能比较函数做测试的时候一个容易犯的错误就是把b.N作为输入的大小

// 测试集的Setup与Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

// 子测试的Setup与Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func TestSplitSetUpANdTearDown(t *testing.T) {
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
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := unit_test.Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

//hfy@HFYdeMac-mini  ~/go/src/awesomeProject/unit_test   master ±✚  go test -v -run=TestSplitSetUpANdTearDown
//write setup code here...
//=== RUN   TestSplitSetUpANdTearDown
//fib_test.go:72: 如有需要在此执行:测试之前的setup
//=== RUN   TestSplitSetUpANdTearDown/more_sep
//fib_test.go:80: 如有需要在此执行:子测试之前的setup
//fib_test.go:82: 如有需要在此执行:子测试之后的teardown
//=== RUN   TestSplitSetUpANdTearDown/leading_sep
//fib_test.go:80: 如有需要在此执行:子测试之前的setup
//fib_test.go:82: 如有需要在此执行:子测试之后的teardown
//=== RUN   TestSplitSetUpANdTearDown/simple
//fib_test.go:80: 如有需要在此执行:子测试之前的setup
//fib_test.go:82: 如有需要在此执行:子测试之后的teardown
//=== RUN   TestSplitSetUpANdTearDown/wrong_sep
//fib_test.go:80: 如有需要在此执行:子测试之前的setup
//fib_test.go:82: 如有需要在此执行:子测试之后的teardown
//=== CONT  TestSplitSetUpANdTearDown
//fib_test.go:74: 如有需要在此执行:测试之后的teardown
//--- PASS: TestSplitSetUpANdTearDown (0.00s)
//--- PASS: TestSplitSetUpANdTearDown/more_sep (0.00s)
//--- PASS: TestSplitSetUpANdTearDown/leading_sep (0.00s)
//--- PASS: TestSplitSetUpANdTearDown/simple (0.00s)
//--- PASS: TestSplitSetUpANdTearDown/wrong_sep (0.00s)
//PASS
//write teardown code here...
//ok      awesomeProject/unit_test        0.397s
