# Command-Line Subcommands
```go
// Một vài công cụ dòng lệnh, giống như công cụ `go` tool
// hoặc `git` có nhiều *lệnh con*, mõi lệnh con có bộ cờ
// riêng của nó. Ví dụ, `go build` và `go get` là hai
// lệnh con khác nhau của công cụ `go`. Package (gói) `flag`
// cho phép chúng ta dễ dàng xác định các lệnh con
// đơn giản có bộ cờ riêng của chúng.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Chúng ta khai báo một lệnh con bằng cách
	// sử dụng hàm `NewFlagSet` và tiếp tục định nghĩa
	// các cờ mới cụ thể cho lệnh con này.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Đối với một lệnh con khác, chúng ta có thể
	// xác định các cờ được hỗ trợ khác nhau.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Lệnh con được mong đợi là đối số đầu tiên
	// cho chương trình.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Kiểm tra lệnh con được kích hoạt.
	switch os.Args[1] {

	// Đối với mỗi lệnh con, chúng tôi phân tích cú pháp
	// các cờ của nó và có quyền truy cập vào các
	// đối số vị trí cuối cùng.
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
```
```sh
$ go build command-line-subcommands.go 

# Đầu tiên kích hoạt lệnh con foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Bây giờ thử bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Nhưng bar sẽ không chấp nhận các cờ của foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Tiếp theo, chúng ta sẽ xem xét biến môi trường,
# một cách phổ biến khác để tham số hóa 
# các chương trình.
```
