package main

//(47)数值
//Go有很多种数据类型，包括字符串类型，整型，浮点型，布尔型等等
import "fmt"

func main() {
	// 字符串可以使用"+"连接
	fmt.Println("go" + "lang")

	//整型和浮点型
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 布尔型的几种操作符
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
