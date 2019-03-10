package main

//(4.7)init函数
import (
	"fmt"
	//引入包
	"github.com/wanglg007/My_GoLang/00_Base/chapter04/utils"
)

var age = test07()

//为了看到全局变量是先被初始化的，我们这里先写函数
func test07() int {
	fmt.Println("test()") //1
	return 90
}

//init函数,通常可以在init函数中完成初始化工作
func init() {
	fmt.Println("init()...") //2
}

func main() {
	fmt.Println("main()...age=", age) //3
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)

}