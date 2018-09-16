package main

//(55)信号处理
//有的时候希望Go能够智能地处理Unix信号。例如希望一个server接收到一个SIGTERM的信号时，能够自动地停止；或者一个命令行工具接收到一个SIGINT信号时，能够停止接收输入。
// 现在来看下如何使用channel来处理信号。
import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	// Go信号通知通过向一个channel发送``os.Signal`来实现。我们将创建一个channel来接受这些通知，同时我们还用
	// 一个channel来在程序可以退出的时候通知我们
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify`在给定的channel上面注册该channel可以接受的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 这个goroutine阻塞等待信号的到来，当信号到来的时候，输出该信号，然后通知程序可以结束
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序将等待接受信号，然后退出
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
