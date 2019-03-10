package main

//(2.9)案例
import (
	"fmt"
)
func main() {

	//案例1：有两个变量，a和b，要求将其进行交换，但是不允许使用中间变量，最终打印结果
	var a int = 10
	var b int = 20

	a = a + b //
	b = a - b // b = a + b - b ==> b = a
	a = a - b // a = a + b - a ==> a = b:
	fmt.Printf("a=%v b=%v", a, b)

	//案例2：求两个数的最大值
	var n1 int = 10
	var n2 int = 40
	var max int
	if n1 > n2 {
		max = n1
	} else {
		max = n2
	}
	fmt.Println("max=", max)

	//求出三个数的最大值思路：先求出两个数的最大值，然后让这个最大值和第三数比较，在取出最大值
	var n3 = 45
	if n3 > max {
		max = n3
	}
	fmt.Println("三个数中最大值是=", max)
}
