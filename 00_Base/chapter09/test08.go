package main

//(9.8)继承多个接口
import (
	"fmt"
)

type BInterface08 interface {
	test01()
}

type CInterface08 interface {
	test02()
}

type AInterface08 interface {
	BInterface08
	CInterface08
	test03()
}

//如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
type Stu08 struct {
}

func (stu Stu08) test01() {

}
func (stu Stu08) test02() {

}
func (stu Stu08) test03() {

}

type T interface {
}

func main() {
	var stu Stu08
	var a AInterface08 = stu
	a.test01()

	var t T = stu //ok
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
