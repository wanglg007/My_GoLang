package main
//(44)时间
//Go提供了对时间和一段时间的支持。
import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 从获取当前时间开始
	now := time.Now()
	p(now)

	// 可以提供年，月，日等来创建一个时间。当然时间总是会和地区联系在一起，也就是时区
	then := time.Date(2010,11,17,20,34,58,651387237,time.UTC)
	p(then)

	// 可以获取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出当天是周几，Monday-Sunday中的一个
	p(then.Weekday())

	// 下面的几个方法判断两个时间的顺序，精确到秒
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub方法返回两个时间的间隔(Duration)
	diff := now.Sub(then)
	p(diff)

	// 可以以不同的单位来计算间隔的大小
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 可以使用Add方法来为时间增加一个间隔,使用负号表示时间向前推移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))

}
