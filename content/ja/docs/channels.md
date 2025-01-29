# Channels
```go
// <em>チャネル（channel）</em>は平行に動くゴルーチンを繋ぐパイプである。
// あるゴルーチンからチャネルに値を送り、別のゴルーチンでその値を受取れるのだ。

package main

import "fmt"

func main() {

	// `make(chan value-type)` で新しいチャネルを作れる。
	// チャネルの型にはそれを通る値の型が入っている。
	messages := make(chan string)

	// `channel <-` と書けばチャネルに値を<em>送信</em>できる。
	// ここでは、新たなゴルーチンから、`"ping"` を先程作ったチャネル `messages` に送っている。
	go func() { messages <- "ping" }()

	// `<-channel` と書けばそのチャネルから値を<em>受信</em>する。
	// ここでは上で送った `"ping"` メッセージを受信し、表示している。
	msg := <-messages
	fmt.Println(msg)
}
```
```sh
# プログラムを起動すると、`"ping"` は一方のゴルーチンから他方のゴルーチンへ、チャネルを通じて届く。
$ go run channels.go 
ping

# デフォルトでは送信側と受信側の両方が準備できるまで、送受信はブロックする。
# この性質によって、プログラムの最後では単に `"ping"` を待つだけで、それ以外の同期が必要なくなっている。
```
