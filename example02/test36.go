package main
//(36) 进程触发
//从Go程序里面触发一个其他的非Go进程来执行
import (
	"os/exec"
	"fmt"
	"io/ioutil"
)

func main() {
	// 从简单的命令开始，这个命令不需要任何参数或者输入，仅仅向stdout输出一些信息。`exec.Command`函数创建了一个代表外部进程的对象
	dateCmd := exec.Command("date")

	// `Output`是另一个运行命令时用来处理信息的函数，这个函数等待命令结束，然后收集命令输出。如果没有错误发生，`dateOut`将保存date的信息
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 下面看一个需要从stdin输入数据的命令，我们将数据输入传给外部进程的stdin，然后从它输出到stdout的运行结果收集信息
	grepCmd := exec.Command("grep", "hello")

	// 这里显式地获取input/output管道，启动进程，向进程写入数据，然后读取输出结果，最后等待进程结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	// 在上面的例子中，我们忽略了错误检测，但是你一样可以使用`if err!=nil`模式来进行处理。另外我们仅仅收集了`StdoutPipe`的结果，同时你也可以用一样的方法来收集
	// `StderrPipe`的结果
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 注意，我们在触发外部命令的时候，需要显式地提供命令和参数信息。另外如果你想用一个命令行字符串触发一个完整的命令，你可以使用bash的-c选项
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
