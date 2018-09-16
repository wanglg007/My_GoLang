package main

//(39)命令行参数
//命令行参数是一种指定程序运行初始参数的常用方式。比如 go run hello.go 使用 run 和 hello.go 参数来执行程序。
import (
	"os"
	"fmt"
)

func main() {
	// `os.Args`提供了对命令行参数的访问，注意该切片的第一个元素是该程序的运行路径，而`os.Args[1:]`则包含了该程序的所有参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 可以使用索引的方式来获取单个参数
	arg := os.Args[3]
	
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
