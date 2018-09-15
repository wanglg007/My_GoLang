package main
//(1)Go状态协程
import (
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

// 本案例将有一个单独的协程拥有这个状态。这样可以保证数据不会被并行访问所破坏。为了读写这个状态，其
// 他的协程将向这个协程发送信息并且相应地接受返回信息。这些`readOp`和`writeOp`结构体封装了这些请求和回复
type readOp struct {
	key int
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}
func main() {
	// 计算执行多少次操作
	var ops int64 = 0

	// reads和writes通道将被其他协程用来从中读取或写入数据
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 拥有state的协程，`state`是一个协程的私有map变量。这个协程不断地`select`通道`reads`和`writes`，当有请求来临的时候进行回复。
	// 一旦有请求，首先执行所请求的操作，然后给`resp`通道发送一个表示请求成功的值。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动了100个协程来向拥有状态的协程请求读数据。每次读操作都需要创建一个`readOp`，然后发送到`reads`通道，然后等待接收请求回复
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 开启10个写协程
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 让协程运行1秒钟
	time.Sleep(time.Second)
	// 最后输出操作数量ops的值
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}


