package main

//(22)打点器
//Timer是让你等待一段时间然后去做一件事情，这件事情只会做一次。而Ticker是让你按照一定的时间间隔
//循环往复地做一件事情，除非你手动停止它。
import (
	"time"
	"fmt"
)

func main() {
	// Ticker使用和Timer相似的机制，同样是使用一个通道来发送数据。这里使用range函数来遍历通道数据，这些数据每隔100毫秒被
	// 发送一次，这样就可以接收到
	ticker := time.NewTicker(time.Microsecond * 100)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Ticker和Timer一样可以被停止。一旦Ticker停止后，通道将不再接收数据，这里我们将在1500毫秒之后停止
	time.Sleep(time.Microsecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
