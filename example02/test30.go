package main
//(30)函数命名返回值
//函数接受参数。在 Go 中，函数可以返回多个“结果参数”，而不仅仅是一个值。
import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
