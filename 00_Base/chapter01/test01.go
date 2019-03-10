package main

//(1.1)golang的变量使用方式
import "fmt"

func main() {
	//(1)指定变量类型，声明后若不赋值，使用默认值
	var i int // int 的默认值是0
	fmt.Println("i=", i)

	//(2)根据值自行判定变量类型(类型推导)
	var num = 10.11
	fmt.Println("num=", num)

	//(3)省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误
	//下面的方式等价 var name string   name = "tom", := 的 :不能省略，否则错误
	name := "tom"
	fmt.Println("name=", name)

}
