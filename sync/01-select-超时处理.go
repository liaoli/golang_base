package sync_demo

import (
	"fmt"
	"time"
)

func test(ch chan int, quit chan bool) {

}

func SelectTimeOutDemo() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() { // 子go 程获取数据
		for {
			fmt.Printf("循环----->")
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):
				quit <- true
				goto lable
				// return
				// runtime.Goexit()
			}
		}
	lable:
	}()
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit // 主go程，阻塞等待 子go程通知，退出
	fmt.Println("finish!!!")
}
