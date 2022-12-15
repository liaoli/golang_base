package sync_demo

import "fmt"

// 死锁1
func deadLockDemo1() {
	ch := make(chan int)
	ch <- 789
	num := <-ch
	fmt.Println("num = ", num)
}

// 死锁2
func deadLockDemo2() {
	ch := make(chan int)
	go func() {
		ch <- 789
	}()
	num := <-ch
	fmt.Println("num = ", num)
}

// 死锁 3
func deadLockDemo3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() { // 子
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}

func DeadLock() {
	deadLockDemo1()
	deadLockDemo2()
	deadLockDemo3()
}
