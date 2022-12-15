package sync_demo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwMutex1 sync.RWMutex // 锁只有一把， 2 个属性 r w

func readGo(in <-chan int, idx int) {
	for {
		rwMutex1.RLock() // 以读模式加锁
		num := <-in
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		rwMutex1.RUnlock() // 以读模式解锁
	}
}

func writeGo(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex1.Lock() // 以写模式加锁
		out <- num
		fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300) // 放大实验现象
		rwMutex1.Unlock()
	}
}

func RwLoakDemo1() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())

	// quit := make(chan bool)			// 用于 关闭主go程的channel
	ch := make(chan int) // 用于 数据传递的 channel

	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}
	for {

	}
}
