package main
//(6)Go Defer
//Defer用来保证一个函数调用会在程序执行的最后被调用。通常用于资源清理工作。

import (
	"os"
	"fmt"
)

// 假设想创建一个文件，然后写入数据，最后关闭文件
func main()  {
	// 在使用createFile得到一个文件对象之后，使用defer来调用关闭文件的方法closeFile，这个方法将在main函数
	// 最后被执行，也就是writeFile完成之后
	f := createFile("d://mm.txt")

	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating:")
	f,err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File)  {
	f.Close()
}