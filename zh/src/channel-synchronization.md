# Channel Synchronization->通道同步
```go
// 我们可以使用通道来同步协程之间的执行状态。
// 这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。
// 如果需要等待多个协程，[WaitGroup](waitgroups) 是一个更好的选择。

package main

import (
	"fmt"
	"time"
)

// 我们将要在协程中运行这个函数。
// `done` 通道将被用于通知其他协程这个函数已经完成工作。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {

	// 运行一个 worker 协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)

	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	<-done
}
```
```sh
$ go run channel-synchronization.go
working...done

# 如果你把 `<- done` 这行代码从程序中移除，
# 程序甚至可能在 `worker` 开始运行前就结束了。
```
