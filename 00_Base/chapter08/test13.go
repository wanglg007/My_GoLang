package main

//(8.13)工厂模式
import (
	"fmt"
	"github.com/wanglg007/My_GoLang/00_Base/chapter08/model"
)

func main() {
	//创建要给Student实例
	// var stu = model.Student{
	// 	Name :"tom",
	// 	Score : 78.9,
	// }

	//定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 98.8)

	fmt.Println(*stu) //&{....}
	fmt.Println("name=", stu.Name, " score=", stu.GetScore())
}
