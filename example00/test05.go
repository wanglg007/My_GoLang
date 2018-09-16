package main

import (
	"fmt"
	"unsafe"
)

/**
(6)指针
支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。
•默认值 nil，没有 NULL 常量。
•操作符 "&" 取变量地址，"*" 透过指针访问⺫⽬目标对象。
•不⽀支持指针运算，不⽀支持 "->" 运算符，直接⽤用 "." 访问⺫⽬目标成员
 */
func main() {
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %v\n", p, p.a) // 直接⽤用指针访问⺫⽬目标对象成员，无须转换。

	//可以在 unsafe.Pointer 和任意类型指针间进⾏行转换
	x := 0x12345678
	p1 := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p1)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}

	//将 Pointer 转换成 uintptr，可变相实现指针运算
	dd := struct {
		s string
		x int
	}{"abc", 100}
	p2 := uintptr(unsafe.Pointer(&dd)) // *struct -> Pointer -> uintptr
	p2 += unsafe.Offsetof(dd.x)        // uintptr + offset

	p3 := unsafe.Pointer(p2) // uintptr -> Pointer
	px := (*int)(p3)         // Pointer -> *int
	*px = 200                // d.x = 200
	fmt.Printf("%#v\n", d)
}
