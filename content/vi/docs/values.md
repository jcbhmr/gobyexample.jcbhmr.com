# Values
```go
// Go có nhiều loại value như strings (chuỗi),
// integers (số nguyên), floats (số thực), booleans (true/false), ...
// Dưới đây là một vài ví dụ cơ bản.

package main

import "fmt"

func main() {

	// Chuỗi có thể liên kết với nhau bằng `+`
	fmt.Println("go" + "lang")

	// Số nguyên và số thực.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans với các toán tử booleans.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```
```sh
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```
