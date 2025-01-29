# Environment Variables
```go
// [環境変数](http://en.wikipedia.org/wiki/Environment_variable)は
// [設定情報を Unix 上で動くプログラムに伝える](http://www.12factor.net/config)一般的な仕組みである。
// 環境変数を読んだり、書いたりしてみよう。

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// キーと値のペアをセットするには `os.Setenv` を使う。
	// キーから値を読み出すには `os.Getenv を使う。
	// もしキーが設定されていなければ空の文字列を返す。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 定義されているキーと値の一覧は `os.Environ` で取得する。
	// この関数は `KEY=value` という形式の文字列からなるスライスを返す。
	// `strings.Split` を使って、このスライスからキーと値を取り出せる。
	// ここではすべてのキーを表示している。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}

}
```
```sh
# プログラムを実行すると、プログラム内で設定した `FOO` の値を読み出せたことがわかる。
# 一方、`BAR` の値が空であることもわかる。
$ go run environment-variables.go
FOO: 1
BAR: 

# 環境におけるキーの一覧は、実行するマシンによって違うはずだ。
TERM_PROGRAM
PATH
SHELL
...

# `BAR` の値を設定してからプログラムを実行すれば、プログラムはその値を読み出せる。
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
```
