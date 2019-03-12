package main

//(9.11)多态

import (
	"fmt"
)

//声明/定义一个接口
type Usb11 interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone11 struct {
	name string
}

//让Phone 实现 Usb接口的方法
func (p Phone11) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone11) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera11 struct {
	name string
}

//让Camera 实现   Usb接口的方法
func (c Camera11) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera11) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb11
	usbArr[0] = Phone11{"vivo"}
	usbArr[1] = Phone11{"小米"}
	usbArr[2] = Camera11{"尼康"}

	fmt.Println(usbArr)

}
