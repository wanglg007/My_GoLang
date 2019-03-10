package main

//(4.3)给函数传入变量地址
import (
	"fmt"
)

// n1 就是 *int 类型
func test03(n1 *int) {
	fmt.Printf("n1的地址 %v\n", &n1)
	*n1 = *n1 + 10
	fmt.Println("test03() n1= ", *n1) // 30
}

func main() {

	num := 20
	fmt.Printf("num的地址=%v\n", &num)
	test03(&num)
	fmt.Println("main() num= ", num) // 30
}
