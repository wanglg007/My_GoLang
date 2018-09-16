package main
//(28)函数定义
//
import "fmt"

// 这个函数计算两个int型输入数据的和，并返回int型的和
func plus(a int, b int) int {
	// Go需要使用return语句显式地返回值
	return a + b
}

func main() {
	// 函数的调用方式很简单 "名称(参数列表)"
	res := plus(1, 2)
	fmt.Println("1+2=", res)
}
