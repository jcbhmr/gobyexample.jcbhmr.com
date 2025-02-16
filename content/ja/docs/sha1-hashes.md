# SHA1 Hashes
```go
// [<em>SHA1 ハッシュ関数</em>](http://en.wikipedia.org/wiki/SHA-1)はバイナリや文字列の短い識別子を計算するためによく使う。
// 例えばバージョン管理システムである [Git](http://git-scm.com/) ではファイルやディレクトリを識別するために SHA1 を使っている。
// それでは、Go で SHA1 を計算する例を見ていこう。

package main

// `crypto/*` パッケージには何種類かのハッシュ関数が含まれている。
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// ハッシュ値を生成するには `sha1.New()`、`sha1.Write(bytes)`、`sha1.Sum([]byte{})` の順で関数を呼ぶ。
	// まずは新たなハッシュ関数を生成する。
	h := sha1.New()

	// `Write` の入力はバイト列である。
	// 文字列 `s` のハッシュ値を計算したいなら、`[]byte(s)` と書いてハッシュ値に変換してやらなければならない。
	h.Write([]byte(s))

	// バイトのスライスとして、最終的なハッシュ値を得る。
	// `Sum` の引数を、これまで入力したバイト列に追記できるが、普通はこれは使わない。
	bs := h.Sum(nil)

	// SHA1 のハッシュ値は、Git がそうしているように、16進記数法で表示することが多い。
	// フォーマット文字列に `%x` と書けば、ハッシュ計算の結果を16進文字列に変換できる。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
```sh
# プログラムはハッシュ値を計算し、人が読みやすい16進記数法で結果を表示する。
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# 上と同様のやり方で他のハッシュ関数を計算することもできる。
# 例えば MD5 ハッシュ関数を使いたければ、
# `crypto/md5` をインポートして `md5.New()` を使えばよい。

# ただし、暗号学的に安全なハッシュ関数を使いたい場合には
# [ハッシュ関数の強度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)について知る必要がある。
```
