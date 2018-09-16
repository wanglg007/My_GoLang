package main

//(45)时间戳
//程序的一个通常需求是计算从Unix起始时间开始到某个时刻的秒数，毫秒数，微秒数等。
import (
	"time"
	"fmt"
)

func main() {
	// 使用Unix和UnixNano来分别获取从Unix起始时间到现在所经过的秒数和微秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意这里没有UnixMillis方法，所以我们需要将微秒手动除以一个数值来获取毫秒
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 反过来，也可以将一个整数秒数或者微秒数转换为对应的时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
