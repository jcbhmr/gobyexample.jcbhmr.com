# Command-Line Flags
```go
// [<em>コマンドラインフラグ</em>](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// はコマンドラインからオプションを指定するためによく使われる。
// 例えば `wc -l` の `-l` はコマンドラインフラグである。

package main

// Go の `flag` パッケージを使ってコマンドラインフラグをパースできる。
// 実際にコマンドラインプログラムを作って、このパッケージを使ってみよう。
import (
	"flag"
	"fmt"
)

func main() {

	// 文字列、整数、真偽値のオプションを受け取るフラグを宣言できる。
	// ここではフラグ `word` を宣言し、そのデフォルト値を `"foo"` とし、フラグの短い説明を与えている。
	// 関数 `flag.String` は（文字列そのものではなく）文字列のポインタを返す。
	// このポインタの使い方は追って説明する。
	wordPtr := flag.String("word", "foo", "a string")

	// `word` と同様のやり方で、フラグ `numb`、`fork` を宣言している。
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// プログラム内の別の場所で宣言した変数を使い、オプションを宣言することもできる。
	// なお、この場合変数のポインタを渡すことに注意する。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// すべてのフラグを宣言したら、`flag.Parse()` を呼んでコマンドラインをパースする。
	flag.Parse()

	// パースしたオプションと、余りの引数を表示する。
	// `*wordPtr` のように参照を剥がしてオプション値を呼んでいることに注意する。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}
```
```sh
# コマンドラインフラグの実験をするときは、まずプログラムをコンパイルし、その後バイナリを直接実行するとよい。
$ go build command-line-flags.go

# ビルドしたプログラムを、すべてのフラグに値を与えて実行する。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# フラグを書かなければ、自動的にデフォルト値を使うことに注意する。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# フラグの後に引数を並べられる。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` パッケージでは、すべてのフラグは単なる引数の前に書かなければならない。
# （さもなければ、フラグはただの引数として解釈されてしまう。）
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# `-h` か `--help` を使うと、自動生成したコマンドラインのヘルプを表示する。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# `flag` パッケージで指定していないフラグをプログラムに渡すと、
# プログラムはエラーメッセージを表示し、この場合もヘルプを表示する。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
```
