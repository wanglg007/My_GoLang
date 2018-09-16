package main
//(26)工作池
//使用gorouotine和channel来实现工作池
import (
	"fmt"
	"time"
)

// 将在worker函数里面运行几个并行实例，这个函数从jobs通道里面接受任务，然后把运行结果发送到results通道。每个job
// 都休眠一会，来模拟一个耗时任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 为了使用工作池，需要发送工作和接受工作的结果，这里定义两个通道，一个jobs，一个results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动3个worker协程，一开始的时候worker阻塞执行，因为jobs通道里面还没有工作任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 然后我们从results里面获得结果
	for a := 1; a <= 9; a++ {
		<-results
	}

}
