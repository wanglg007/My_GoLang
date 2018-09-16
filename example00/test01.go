package main

/**
(2)类型-常量:
常量值必须是编译期可确定的数字、字符串、布尔值
 */
import (
	"fmt"
	"unsafe"
)

const x, y int = 1, 2     // 多常量初始化
const s = "Hello, World!" // 类型推断
const ( // 常量组
	a, b      = 10, 100
	c    bool = false

	//常量值还可以是 len、cap、unsafe.Sizeof 等编译期可确定结果的函数返回值
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(b)

	//如果常量类型⾜足以存储初始化值，那么不会引发溢出错误
	a2 byte = 100  // int to byte
	//b2 int  = 1e20 // float64 to int, overflows
)

const ( //关键字 iota 定义常量组中从 0 开始按⾏行计数的⾃自增枚举值
	Sunday    = iota // 0
	Monday           // 1，通常省略后续⾏行表达式。
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // 与 KB 表达式相同，但 iota = 2
	GB
	TB
)

//可通过⾃自定义类型来实现枚举类型限制
type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test1(c Color) {}

func main() {
	const x = "xxx" // 未使⽤用局部常量不会引发编译错误。
	fmt.Println(x)

	fmt.Println(x, y, s, a, b, c)

	fmt.Println(a2)

	fmt.Println(Monday)

	c2 := Black
	test1(c2)
	x2 := 1
	//test(x2) 			// Error: cannot use x (type int) as type Color in function argument

	test1(1) // 常量会被编译器⾃自动转换。
	fmt.Println(x2)

}
