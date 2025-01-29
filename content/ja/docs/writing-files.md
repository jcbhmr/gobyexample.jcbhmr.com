# Writing Files
```go
// Go でファイルに書き込むには、前に見たファイルを読み出すのと似たパターンを使う。

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// まずは文字列（またはバイト列）を単にファイルに書き込む。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// より細かな制御をするために、ファイルを書き込み用に開く。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// ファイルを開いた直後に `Close` を `defer` するのはイディオムである。
	defer f.Close()

	// `Write` を使ってバイトのスライスを書き込める。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` という関数もあり、こちらは文字列を書き込める。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// `Sync` を使って書き込みをストレージにフラッシュする。
	f.Sync()

	// `bufio` を使ってバッファ付きのリーダーを作る例を前に紹介した。
	// 同様に、バッファ付きのライターも作れる。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// `Flush` を使って、ライターのバッファされている操作をすべてフラッシュする。
	w.Flush()

}
```
```sh
# ファイルに書き込みをするプログラムを実行してみる。
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# そして、書き込み対象のファイルの内容を確認する。
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 続いて、ここまで見てきたファイル IO のアイデアを、
# ストリームである `stdin` と `stdout` に適用する例を紹介する。
```
