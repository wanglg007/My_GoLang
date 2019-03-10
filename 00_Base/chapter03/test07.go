package main

//(3.7)goto
import (
	"fmt"
)

func main() {

	var n int = 30
	//演示goto的使用
	fmt.Println("ok1")
	if n > 20 {
		goto label1
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
label1:
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")

	var n1 int = 30
	//演示return的使用
	fmt.Println("ok1")
	if n1 > 20 {
		return
	}
	fmt.Println("ok2")
	fmt.Println("ok3")
	fmt.Println("ok4")
	fmt.Println("ok5")
	fmt.Println("ok6")
	fmt.Println("ok7")
}
