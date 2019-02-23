package main

//(1.5)演示golang中整数类型使用
import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 1
	fmt.Println("i=", i)

	//测试一下int8的范围 -128~127
	var j int8 = 127
	fmt.Println("j=", j)

	//测试一下uint8的范围(0-255)
	var k uint16 = 255
	fmt.Println("k=", k)

	//int , uint , rune , byte的使用
	var a int = 8900
	fmt.Println("a=", a)
	var b uint = 1
	var c byte = 255
	fmt.Println("b=", b, "c=", c)

	//整型的使用细节
	var n1 = 100
	//介绍一下如何查看某个变量的数据类型
	fmt.Printf("n1 的 类型 %T \n", n1)

	//在程序查看某个变量的占用字节大小和数据类型 （使用较多）
	var n2 int64 = 10
	//unsafe.Sizeof(n1) 是unsafe包的一个函数，可以返回n1变量占用的字节数
	fmt.Printf("n2 的 类型 %T  n2占用的字节数是 %d ", n2, unsafe.Sizeof(n2))

}
