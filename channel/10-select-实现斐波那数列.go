package channel_demo

import (
	"awesomeProject/channel/data_stream"
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Print(num, " ")
		case <-quit:
			//return
			runtime.Goexit() //等效于 return
		}
	}
}

func ChannelDemo10() {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit) // 子go 程 打印fibonacci数列

	x, y := 1, 1
	for i := 0; i < 40; i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}

func ChannelDemo() {
	//ChannelDemo1()
	//ChannelDemo2()
	//ChannelDemo3()
	//ChannelDemo4()
	//ChannelDemo5()
	//ChannelDemo6()
	//ChannelDemo7()
	//ChannelDemo9()
	//ChannelDemo10()
	//ChannelDemo11()
	//ChannelDemo12()
	//ChannelDemo13()
	//PlayBallDemo()
	data_stream.DataStreamDemo()
}
