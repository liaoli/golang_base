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

//其实在 Go 程序启动时，Go 程序就会为 main 函数创建一个默认的 goroutine 。
//在上面的代码中我们在 main 函数中使用 go 关键字创建了另外一个 goroutine 去执行 hello 函数，
//而此时 main goroutine 还在继续往下执行，我们的程序中此时存在两个并发执行的 goroutine。
//当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，所有由 main goroutine 创建的 goroutine 也会一同退出。
//也就是说我们的 main 函数退出太快，另外一个 goroutine 中的函数还未执行完程序就退出了，导致未打印出“hello”。
//
//main goroutine 就像是《权利的游戏》中的夜王，其他的 goroutine 都是夜王转化出的异鬼，夜王一死它转化的那些异鬼也就全部GG了。
//
//所以我们要想办法让 main 函数‘“等一等”将在另一个 goroutine 中运行的 hello 函数。
//其中最简单粗暴的方式就是在 main 函数中“time.Sleep”一秒钟了
//（这里的1秒钟只是我们为了保证新的 goroutine 能够被正常创建和执行而设置的一个值）。

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

//sync.WaitGroup
//在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。 sync.WaitGroup有以下几个方法：
//
//方法名	功能
//func (wg * WaitGroup) Add(delta int)	计数器+delta
//(wg *WaitGroup) Done()	计数器-1
//(wg *WaitGroup) Wait()	阻塞直到计数器变为0
//sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了 N 个并发任务时，就将计数器值增加N。
//每个任务完成时通过调用 Done 方法将计数器减1。通过调用 Wait 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。

func waitGroupDemo() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			print(i)
		}(i)
	}

	wg.Wait()
}

func print(i int) {
	defer wg.Done()
	fmt.Println(i)
}

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}
func waitGroupDemo2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

//Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
//
//如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
//
//Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。
//
//channel类型
//channel是 Go 语言中一种特有的类型。声明通道类型变量的格式如下：
//
//var 变量名称 chan 元素类型
//其中：
//
//chan：是关键字
//元素类型：是指通道中传递元素的类型
//举几个例子：
//
//var ch1 chan int   // 声明一个传递整型的通道
//var ch2 chan bool  // 声明一个传递布尔型的通道
//var ch3 chan []int // 声明一个传递int切片的通道
//channel零值
//未初始化的通道类型变量其默认零值是nil。
//
//var ch chan int
//fmt.Println(ch) // <nil>
//初始化channel
//声明的通道类型变量需要使用内置的make函数初始化之后才能使用。具体格式如下：
//
//make(chan 元素类型, [缓冲大小])
//其中：
//
//channel的缓冲大小是可选的。
//举几个例子：
//
//ch4 := make(chan int)
//ch5 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道
//channel操作
//通道共有发送（send）、接收(receive）和关闭（close）三种操作。而发送和接收操作都使用<-符号。
//
//现在我们先使用以下语句定义一个通道：
//
//ch := make(chan int)
//发送
//将一个值发送到通道中。
//
//ch <- 10 // 把10发送到ch中
//接收
//从一个通道中接收值。
//
//x := <- ch // 从ch中接收值并赋值给变量x
//<-ch       // 从ch中接收值，忽略结果
//关闭
//我们通过调用内置的close函数来关闭通道。
//
//close(ch)
//注意：一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
//
//关闭后的通道有以下特点：
//
//对一个关闭的通道再发送值就会导致 panic。
//对一个关闭的通道进行接收会一直获取值直到通道为空。
//对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//关闭一个已经关闭的通道会导致 panic。

//无缓冲的通道
//无缓冲的通道又称为阻塞的通道。我们来看一下如下代码片段。

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
	wg.Done()
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
	wg.Add(1)
	go recv(ch)
	ch <- 10
	fmt.Println("发送成功")
	wg.Wait()
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
	ch := make(chan int, 6)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
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

//并发数据安全问题

var (
	x int64
)

// add 对全局变量x执行5000次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

//当有多个goroutine 操作同一个数据的时候就会出现数据安全问题
//可以通过锁来解决
func syncDemo() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}

//互斥锁 sync.Mutex
var m sync.Mutex // 互斥锁 确保同一时间只有一个goroutine在操作数据
func addWithMutexLock() {
	for i := 0; i < 5000; i++ {
		m.Lock() //申请锁
		x = x + 1
		m.Unlock() //使用完成释放锁
	}
	wg.Done()
}

//使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；
//当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，
//唤醒的策略是随机的。
func syncMutexLockDemo() {
	wg.Add(2)

	go addWithMutexLock()
	go addWithMutexLock()

	wg.Wait()
	fmt.Println(x)
}

//读写互斥锁
//互斥锁是完全互斥的，但是实际上有很多场景是读多写少的，当我们并发的去读取一个资源而不涉及资源修改的时候是没有必要加互斥锁的，
//这种场景下使用读写锁是更好的一种选择。读写锁在 Go 语言中使用sync包中的RWMutex类型
//sync.RWMutex提供了以下5个方法。
//
//方法名	功能
//func (rw *RWMutex) Lock()	获取写锁
//func (rw *RWMutex) Unlock()	释放写锁
//func (rw *RWMutex) RLock()	获取读锁
//func (rw *RWMutex) RUnlock()	释放读锁
//func (rw *RWMutex) RLocker() Locker	返回一个实现Locker接口的读写锁
//读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
//而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。
//
//下面我们使用代码构造一个读多写少的场景，然后分别使用互斥锁和读写锁查看它们的性能差异。

func goroutinesDemo() {
	//demo1()
	//waitGroupDemo()
	//waitGroupDemo2()
	channelDemo1()
	//channelDemo2()
	//channelDemo3()
	//channelDemo4()
	//channelDemo5()
	//channelDemo6()
	//channelDemo7()
	//channelDemo8()
	//syncDemo()
	//syncMutexLockDemo()
}
