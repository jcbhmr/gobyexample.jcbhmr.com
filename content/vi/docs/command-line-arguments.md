# Command-Line Arguments
```go
// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// Là một cách để tham số hóa việc chạy chương trình.
// Ví dụ, `go run hello.go` sử dụng tham số `run` và
// `hello.go` cho chương trình `go`.

package main

import (
	"fmt"
	"os"
)

func main() {
	// `os.Args` cung cấp truy cập đến tham số
	// command-line nguyên bản. Lưu ý rằng, giá trị đầu
	// tiên trong slice này là đường dẫn chương trình,
	// và `os.Args[1:]` giữ tham số của chương trình.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Bạn có thể lấy từng thành phần riêng biệt với
	// index thông thường.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
```
```sh
# Để thử nghiệm với tham số command-line, tốt nhất là 
# build binary với `go build` trước.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
# Sau đó, chúng ta sẽ đi sâu hơn vào xử lý command-line
# với flag.
```
