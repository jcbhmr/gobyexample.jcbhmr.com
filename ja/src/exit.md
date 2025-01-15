# Exit
```go
// `os.Exit` を使うと引数に渡したステータスで直ちにプログラムを終了する。

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Exit` を使うと `defer` は実行 _されない_ 。
	// そのためこの `fmt.Println` は呼ばれない。
	defer fmt.Println("!")

	// ステータス 3 で終了する。
	os.Exit(3)
}

// C などと違って、Go は `main` 関数が返す整数型の値を使って終了ステータスを表すことはない。
// そのため、もしゼロでないステータスを返したければ、`os.Exit` を使わなければならない。
```
```sh
# `exit.go` を `go run` で実行すると、
# プログラム終了は `go ` に補足され、画面に表示される。
$ go run exit.go
exit status 3

# バイナリをビルドしてから実行する場合は、ターミナル上でステータスを確認できる。
$ go build exit.go
$ ./exit
$ echo $?
3

# なお、プログラムによって `!` が表示されることはない。
```
