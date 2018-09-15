package main
//(12)Go Panic
//Panic表示的意思就是有些意想不到的错误发生。通常用来表示程序正常运行过程中不应该出现的，或者没有处理好的错误。
import "os"

func main() {
	// 我们使用panic来检查预期不到的错误
	panic("a problem")

	// Panic的通常使用方法就是如果一个函数返回一个我们不知道怎么处理的错误的时候，直接终止执行。
	_, err := os.Create("d://mm.txt")
	if err != nil {
		panic(err)
	}
}
