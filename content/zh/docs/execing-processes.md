# Exec'ing Processes->执行进程
```go
// 在前面的例子中，我们了解了[生成外部进程](spawning-processes)的知识，
// 当我们需要在运行的 Go 流程中访问的外部流程时，便可以执行此操作。
// 但是有时候，我们只想用其它（也许是非 Go）的进程，来完全替代当前的 Go 进程。
// 这时，我们可以使用经典的 <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a> 函数的 Go 的实现。

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 在这个例子中，我们将执行 `ls` 命令。
	// Go 要求我们提供想要执行的可执行文件的绝对路径，
	// 所以我们将使用 `exec.LookPath` 找到它（应该是 `/bin/ls`）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` 需要的参数是切片的形式的（不是放在一起的一个大字符串）。
	// 我们给 `ls` 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` 同样需要使用[环境变量](environment-variables)。
	// 这里我们仅提供当前的环境变量。
	env := os.Environ()

	// 这里是真正的 `syscall.Exec` 调用。
	// 如果这个调用成功，那么我们的进程将在这里结束，并被 `/bin/ls -a -l -h` 进程代替。
	// 如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
```
```sh
# 当我们运行程序时，它会替换为 `ls`。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# 注意 Go 没有提供 Unix 经典的  `fork` 函数。
# 一般来说，这没有问题，因为启动协程、生成进程和执行进程，
# 已经涵盖了 fork 的大多数使用场景。
```
