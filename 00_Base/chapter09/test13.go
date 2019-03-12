package main

//(9.13)类型断言案例
import (
	"fmt"
)

//声明/定义一个接口
type Usb13 interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

type Phone13 struct {
	name string
}

//让Phone 实现 Usb接口的方法
func (p Phone13) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone13) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (p Phone13) Call() {
	fmt.Println("手机 在打电话..")
}

type Camera13 struct {
	name string
}

//让Camera 实现   Usb接口的方法
func (c Camera13) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera13) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer13 struct {
}

func (computer Computer13) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法
	//类型断言..[注意体会!!!]
	if phone, ok := usb.(Phone13); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb13
	usbArr[0] = Phone13{"vivo"}
	usbArr[1] = Phone13{"小米"}
	usbArr[2] = Camera13{"尼康"}

	//遍历usbArr
	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb 接口声明的方法外，还需要调用Phone 特有方法 call. =》类型断言
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	//fmt.Println(usbArr)
}
