# Hello World
```go
// はじめのプログラムは "hello world" を表示するものだ。
// ソースコードのここにあるものですべてだ（ソースコードの右上にある Gopher 君をクリックすると、ブラウザ上でプログラムを実行できる）。
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
```sh
# プログラムを実行するには、`hello-world.go` にコードを書いて、
# `go run` を実行すればよい。
$ go run hello-world.go
hello world

# プログラムをビルドしてバイナリにしたいこともある。
# このためには `go build` を実行する。
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# そうすると、バイナリを直接実行できる。
$ ./hello-world
hello world

# ここまでで、Go のプログラムをビルド・実行できるようになった。
# それでは、言語自体について学んでいこう。
```
