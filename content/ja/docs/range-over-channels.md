# Range over Channels
```go
// [前の例](range)で、`for` と `range` を使って基本的なデータ構造上の繰り返しが書けることを紹介した。
// この記法を使って、チャネルから値を繰り返し受信することもできる。

package main

import "fmt"

func main() {

	// この例では queue チャネルに送信された2つの値を処理する。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// この `range` は `queue` から要素を繰り返し取得する。
	// 既に上でチャネルを `close` しているので、この繰り返しは2つ要素を受信すると終わる。
	for elem := range queue {
		fmt.Println(elem)
	}
}
```
```sh
$ go run range-over-channels.go
one
two

# この例を見ると、空でないチャネルを閉じることもできて、
# その場合にもチャネルに残っている値を受信できることがわかる。
```
