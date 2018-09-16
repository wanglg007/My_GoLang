package main

//(53)通道选择Select
//Go的select关键字可以让你同时等待多个通道操作，将协程（goroutine），通道（channel）和select结合起来构成了Go的一个强大特性。
import (
	"time"
	"fmt"
)

func main() {
	// 本例中，我们从两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)

	// 为了模拟并行协程的阻塞操作，我们让每个通道在一段时间后再写入一个值
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// 我们使用select来等待这两个通道的值，然后输出
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
