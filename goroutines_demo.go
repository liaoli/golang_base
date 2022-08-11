package main

import (
	"fmt"
	"sync"
	"time"
)

/**
*@author: 廖理
*@date:2022/8/10
**/
//学习播客地址：https://www.liwenzhou.com/posts/Go/concurrence/
var wg sync.WaitGroup

func demo1() {
	go spinner(100 * time.Millisecond)
	const n = 44
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func channelDemo() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}

//fatal error: all goroutines are asleep - deadlock!

func channelDemo1() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}

//deadlock表示我们程序中的 goroutine 都被挂起导致程序死锁了。为什么会出现deadlock错误呢？
//
//因为我们使用ch := make(chan int)创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段。
//同理，如果对一个无缓冲通道执行接收操作时，没有任何向通道中发送值的操作那么也会导致接收操作阻塞。就像田径比赛中的4x100接力赛，想要完成交棒必须
//有一个能够接棒的运动员，否则只能等待。简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功。

//上面的代码会阻塞在ch <- 10这一行代码形成死锁，那如何解决这个问题呢？

func channelDemo2() {
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//首先无缓冲通道ch上的发送操作会阻塞，直到另一个 goroutine 在该通道上执行接收操作，这时数字10才能发送成功，
//两个 goroutine 将继续执行。相反，如果接收操作先执行，接收方所在的 goroutine 将阻塞，
//直到 main goroutine 中向该通道发送数字10。
//使用无缓冲通道进行通信将导致发送和接收的 goroutine 同步化。因此，无缓冲通道也被称为同步通道。

//只要通道的容量大于零，那么该通道就属于有缓冲的通道，通道的容量表示通道中最大能存放的元素数量。
//当通道内已有元素数达到最大容量后，再向通道执行发送操作就会阻塞，除非有从通道执行接收操作。
//就像你小区的快递柜只有那么个多格子，格子满了就装不下了，就阻塞了，等到别人取走一个快递员就能往里面放一个。
//
//我们可以使用内置的len函数获取通道内元素的数量，使用cap函数获取通道的容量，虽然我们很少会这么做。

func channelDemo3() {
	ch := make(chan int, 1)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
}

//注意：一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
//它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
//
//关闭后的通道有以下特点：
//
//对一个关闭的通道再发送值就会导致 panic。
//对一个关闭的通道进行接收会一直获取值直到通道为空。
//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//关闭一个已经关闭的通道会导致 panic。

//当向通道中发送完数据时，我们可以通过close函数来关闭通道。
//当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。
//通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。
//那我们如何判断一个通道是否被关闭了呢？
//
//对一个通道执行接收操作时支持使用如下多返回值模式。
//
//value, ok := <- ch
//其中：
//
//value：从通道中取出的值，如果通道被关闭则返回对应类型的零值。
//ok：通道ch关闭时返回 false，否则返回 true。
//下面代码片段中的f2函数会循环从通道ch中接收所有值，直到通道被关闭后退出。

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func channelDemo4() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}

func f3(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

//通常我们会选择使用for range循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
//上面那个示例我们使用for range改写后会很简洁。
//注意：目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。
//不能简单的通过len(ch)操作来判断通道是否被关闭。
func channelDemo5() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f3(ch)
}

//在某些场景下我们可能会将通道作为参数在多个任务函数间进行传递，通常我们会选择在不同的任务函数中对通道的使用进行限制，
//比如限制通道在某个函数中只能执行发送或只能执行接收操作。想象一下，我们现在有Producer和Consumer两个函数，其中Producer函数会返回一个通道，
//并且会持续将符合条件的数据发送至该通道，并在发送完成后将该通道关闭。而Consumer函数的任务是从通道中接收值进行计算，
//这两个函数之间通过Processer函数返回的通道进行通信。完整的示例代码如下。

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

//生产者消费者
func channelDemo6() {
	ch := Producer()

	res := Consumer(ch)
	fmt.Println(res) // 25

}

//从上面的示例代码中可以看出正常情况下Consumer函数中只会对通道进行接收操作，但是这不代表不可以在Consumer函数中对通道进行发送操作。
//作为Producer函数的提供者，我们在返回通道的时候可能只希望调用方拿到返回的通道后只能对其进行接收操作。
//但是我们没有办法阻止在Consumer函数中对通道进行发送操作。
//
//Go语言中提供了单向通道来处理这种需要限制通道只能进行某种操作的情况。
//
//<- chan int // 只接收通道，只能接收不能发送
//chan <- int // 只发送通道，只能发送不能接收
//其中，箭头<-和关键字chan的相对位置表明了当前通道允许的操作，这种限制将在编译阶段进行检测。
//另外对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。
//
//我们使用单向通道将上面的示例代码进行如下改造。

func Producer1() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer1 从通道中接收数据进行计算
func Consumer1(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

//生产者消费者
func channelDemo7() {
	ch := Producer1()

	res := Consumer1(ch)
	fmt.Println(res) // 25

}

//这一次，Producer函数返回的是一个只接收通道，这就从代码层面限制了该函数返回的通道只能进行接收操作，保证了数据安全。
//很多读者看到这个示例可能会觉着这样的限制是多余的，但是试想一下如果Producer函数可以在其他地方被其他人调用，
//你该如何限制他人不对该通道执行发送操作呢？并且返回限制操作的单向通道也会让代码语义更清晰、更易读。
//在函数传参及任何赋值操作中全向通道（正常通道）可以转换为单向通道，但是无法反向转换。

//select多路复用

//Select 语句具有以下特点。
//
//可处理一个或多个 channel 的发送/接收操作。
//如果多个 case 同时满足，select 会随机选择一个执行。
//对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。
//下面的示例代码能够在终端打印出10以内的奇数，我们借助这个代码片段来看一下 select 的具体使用。

//第一次循环时 i = 1，select 语句中包含两个 case 分支，此时由于通道中没有值可以接收，
//所以x := <-ch 这个 case 分支不满足，而ch <- i这个分支可以执行，会把1发送到通道中，结束本次 for 循环；
//第二次 for 循环时，i = 2，由于通道缓冲区已满，所以ch <- i这个分支不满足，而x := <-ch这个分支可以执行，
//从通道接收值1并赋值给变量 x ，所以会在终端打印出 1；
//后续的 for 循环以此类推会依次打印出3、5、7、9。
func channelDemo8() {
	ch := make(chan int, 2)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
func goroutinesDemo() {
	channelDemo8()
}