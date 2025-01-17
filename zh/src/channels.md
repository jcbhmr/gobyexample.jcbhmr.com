# Channels->通道
```go
// _通道(channels)_ 是连接多个协程的管道。
// 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

package main

import "fmt"

func main() {

	// 使用 `make(chan val-type)` 创建一个新的通道。
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)

	// 使用 `channel <-` 语法 _发送_ 一个新的值到通道中。
	// 这里我们在一个新的协程中发送 `"ping"` 到上面创建的 `messages` 通道中。
	go func() { messages <- "ping" }()

	// 使用 `<-channel` 语法从通道中 _接收_ 一个值。
	// 这里我们会收到在上面发送的 `"ping"` 消息并将其打印出来。
	msg := <-messages
	fmt.Println(msg)
}
```
```sh
# 我们运行程序时，通过通道，
# 成功的将消息 `"ping"` 从一个协程传送到了另一个协程中。
$ go run channels.go
ping

# 默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。
# 这个特性允许我们，不使用任何其它的同步操作，
# 就可以在程序结尾处等待消息 `"ping"`。
```
