package main

import "fmt"

/**
(4)类型-引用及转换
引⽤用类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。
内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译成具体的创建函数，由其分配内存和初始化成员结构，返回对象⽽非指针。
 */
func main() {
	a := []int{0, 0, 0} // 提供初始化表达式。
	a[1] = 10
	b := make([]int, 3) // make slice
	b[1] = 10
	c := new([]int)
	//c[1] = 10 			// Error: invalid operation: c[1] (index of type *[]int)
	fmt.Println(a,b,c)

	//不支持隐式类型转换，即便是从窄向宽转换也不⾏
	var b1 byte = 100
	// var n int = b 		// Error: cannot use b (type byte) as type int in assignment
	var n1 int = int(b1) 	// 显式转换
	fmt.Println(b1,n1)

	//同样不能将其他类型当 bool 值使⽤
	a1 := 100
	//if a1 { 					// Error: non-bool a (type int) used as if condition
	//	println("true")
	//}
	fmt.Println(a1)
}
