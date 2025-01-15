# Goroutines
```go
// Một _goroutine_ một luông thực thi lightweight (nhẹ).

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Giả sử chúng ta gọi một hàm `f(s)`. Dưới đây là cách
	// chúng ta gọi hàm đó theo cách thông thường, thực thi
	// đồng bộ.
	f("direct")

	// Để gọi hàm này trong một goroutine, sử dụng cú pháp
	// `go f(s)`. Goroutine mới này sẽ thực thi song song
	// với gouroutine gọi hàm gốc.
	go f("goroutine")

	// Bạn cũng có thể bắt đầu một goroutine cho một lần
	// gọi hàm ẩn danh.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Hai lần gọi hàm của chúng ta đang được thực thi bất đồng bộ
	// trong các goroutine riêng biệt. Đợi chúng thực thi xong
	// (để có một cách tiếp cận đáng tin cậy hơn, ta có thể sử dụng
	// một [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}
```
```sh
# Khi chúng ta chạy chương trình này, chúng ta sẽ thấy
# đầu tiên là kết quả của lần gọi đồng bộ, sau đó là
# kết quả của hai goroutines. Kết quả của các 
# goroutines' có thể được xen kẽ lẫn nhau vì các 
# goroutine đang được thực thi song song bởi 
# Go runtime.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Tiếp theo, chúng ta sẽ xem xét một phần bổ sung 
# cho goroutine trong các concurrent Go program 
#(chương trình Go  đồng thời): channels (kênh).
```
