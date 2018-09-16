package main

//(37)进程执行
//从Go进程里面可以访问外部进程的信息。但有的时候，仅仅希望执行一个外部进程来替代当前的Go进程。这时需要使用Go提供的 exec 函数。
import (
	"os/exec"
	"os"
	"syscall"
)

func main() {
	// 本例中，使用`ls`来演示。Go需要一个该命令的完整路径，所以我们使用`exec.LookPath`函数来找到它
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec`函数需要一个切片参数，给ls命令一些常见的参数。注意，第一个参数必须是程序名称
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec`还需要一些环境变量，这里我们提供当前的系统环境
	env := os.Environ()

	// 这里是`os.Exec`调用。如果一切顺利，我们的原进程将终止，然后启动一个新的ls进程。如果有错误发生，我们将获得一个返回值
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
