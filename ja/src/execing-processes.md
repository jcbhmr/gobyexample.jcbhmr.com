# Exec'ing Processes
```go
// 前の例で[外部プロセスを作る](spawning-processes.html)方法を紹介した。
// 実行中の Go プロセスから外部プロセスにアクセスするときはこの機能を使う。
// しかし、Go のプロセスを別の（Go ではないかもしれない）プロセスに置き換えたい場面もある。
// このようなときは昔からある <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a> の機能を、Go で実装したものを使う。

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// この例では `ls` を exec する。
	// 実行したいバイナリの絶対パスを Go は必要とするので、
	// `exec.LookPath` を使う（おそらく `/bin/ls` だろう）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` の引数はスライスで表現する（ひとつの大きな文字列ではない）。
	// ここでは `ls` によく使う引数を渡してみる。
	// なお、最初の引数はプログラム名であることに注意する。
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` には[環境変数](environment-variables.html)も渡す必要がある。
	// ここでは、現在の環境変数をそのまま渡す。
	env := os.Environ()

	// `syscall.Exec` を呼ぶ。
	// 呼び出しが成功すると、このプロセスはここで終わり、`/bin/ls -a -l -h` を実行するプロセスに置き換わる。
	// もしエラーがあれば、それを表す値が返ってくる。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
```
```sh
# プログラムを実行すると、そのプログラムは `ls` に置き換わる。
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Go は Unix の古典的な `fork` 関数を提供していないことに注意する。
# しかし、普通はそれで問題はない。
# なぜなら、ゴルーチンを起動したり、プロセスを spawn、exec したりできれば、
# `fork` のユースケースの大部分をカバーできるからである。
```
