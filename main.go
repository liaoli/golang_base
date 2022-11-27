package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var a int64

func main() {

	//scope()
	//
	//anonymousFunctionDemo()
	//
	//deferDemo1()

	//deferDemo2()

	//deferDemo3()

	//oopDemo()

	//referencePointer()
	//testPeptide()

	//testEmptyInterface()
	//errorDemo()
	//testPanic()

	//testRecover()
	//stringDemo()
	//deleteSliceItem()

	//r:=rand.Perm(10) //获取随机数的工具类
	//
	//fmt.Print(r)

	//structDemo()
	//FileDemo()
	//goroutinesDemo()
	//lockDemo()
	//syncSysMap()
	//syncMapDemo()

	//date := "20221001"
	//
	//t, err := time.Parse( "20060102",date )
	//
	//if err == nil {
	//	println(t.String())
	//}

	//tcp.StartServer3()
	//udp.StartServer()
	//myhttp.StartServer()
	//res := math.Pow(1.1, 8)*40
	//
	//fmt.Println(res)

	//DataType()
	//ArrayDemo()
	//SliceDemo()
	//mapDemo()
	//FunctionDemo()
	//deferTestDemo()
	//testRecover()
	//fmtDemo()
	//TypeDemo()
	//StructDemo()

	ts := "2022年10月1日 - 2022年10月31日"

	parseTime(ts)

}

func parseTime(ts string) {
	arr1 := strings.Split(ts, " - ")

	for _, v := range arr1 {

		year, ms, err := parseYMD(v, "年")

		if err != nil {

		}
		month, ds, err := parseYMD(ms, "月")
		if err != nil {

		}
		day, _, err := parseYMD(ds, "日")

		// 获取北京时间所在的东八区时区对象

		loc, err := time.LoadLocation("Asia/Shanghai")
		mytime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
		fmt.Println()
		formatMyTime := mytime.Format("2006-01-02 15:04:05")

		fmt.Println(formatMyTime)

		timestamp := mytime.Unix()

		fmt.Printf("时间戳：%d\n", timestamp)

	}
}

func parseYMD(v, sp string) (int, string, error) {
	arr2 := strings.Split(v, sp)

	year, err := strconv.Atoi(arr2[0])

	if err != nil {
		//时间格式错误
		return 0, "", err
	}
	fmt.Printf("%d%s", year, sp)

	return year, arr2[1], err
}
