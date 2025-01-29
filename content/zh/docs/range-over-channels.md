# Range over Channels->通道遍历
```go
// 在[前面](range)的例子中，
// 我们讲过 `for` 和 `range` 为基本的数据结构提供了迭代的功能。
// 我们也可以使用这个语法来遍历的从通道中取值。

package main

import "fmt"

func main() {

	// 我们将遍历在 `queue` 通道中的两个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// `range` 迭代从 `queue` 中得到每个值。
	// 因为我们在前面 `close` 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。
	for elem := range queue {
		fmt.Println(elem)
	}
}
```
```sh
$ go run range-over-channels.go
one
two

# 这个例子也让我们看到，一个非空的通道也是可以关闭的，
# 并且，通道中剩下的值仍然可以被接收到。
```
