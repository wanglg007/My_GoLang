package main
//(13)Go range函数
//range函数是个神奇而有趣的内置函数，你可以使用它来遍历数组，切片和字典。
//当用于遍历数组和切片的时候，range函数返回索引和元素；
//当用于遍历字典的时候，range函数返回字典的键和值。
import "fmt"

// 这里使用range来计算一个切片的所有元素和这种方法对数组也适用
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range 用来遍历数组和切片的时候返回索引和元素值。如果不要关心索引可以使用一个下划线(_)来忽略这个返回值
	// 当然有的时候也需要这个索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 使用range来遍历字典的时候，返回键值对。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range函数用来遍历字符串时，返回Unicode代码点。第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
