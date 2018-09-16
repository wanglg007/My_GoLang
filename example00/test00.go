package main

import "fmt"

/**
(1)类型-变量:
Go是静态类型语言，不能在运⾏期改变量类型。
 */

func test() (int, string) {
	return 1, "abc"
}

func main() {
	//(1)使⽤用关键字 var 定义变量，自动初始化为零值。如果提供初始化值，可省略变量类型，由编译器⾃自动推断。
	var x0 int
	var x1 float32 = 1.6
	var x2 = "abc"
	fmt.Println(x0, x1, x2)

	//(2)在函数内部，可⽤用更简略的 ":=" 方式定义变量。
	x3 := 123
	//一次定义多个变量
	var x4, x5 = "abc", 123
	fmt.Println(x3, x4, x5)

	//(3)多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 // (i = 0) -> (i = 2), (data[0] = 100)

	//(4)特殊只写变量 "_"，⽤用于忽略值占位
	_, y := test()
	println(y)
}
