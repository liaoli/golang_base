package channel_demo

import (
	"fmt"
	"sync"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("生产：", i*i)
		out <- i * i
	}
	close(out)
	w.Done()
}

func consumer(in <-chan int, index int) {
	for num := range in {
		fmt.Printf("消费者%d拿到：%d\n", index, num)
		//time.Sleep(time.Second)
	}
	w.Done()
}

var w sync.WaitGroup

func ChannelDemo6() {
	ch := make(chan int, 6)
	w.Add(3)
	go producer(ch)    // 子go程 生产者
	go consumer(ch, 1) // 主go程 消费

	go consumer(ch, 2) // 主go程 消费

	w.Wait()
}
