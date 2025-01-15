# Channels
```go
// _Channels_ là những ống mà kết nối những goroutine
// đồng thời. Bạn có thể gửi những giá trị vào trong
// những channel từ một goroutine và nhận những giá trị
// trong một goroutine khác.

package main

import "fmt"

func main() {
	// Tạo một channel mới với `make(chan val-type)`.
	// Những channel được nhập theo các giá trị mà
	// chúng truyền tải.
	messages := make(chan string)

	// _Gửi_ một giá trị vào trong một channel sử dụng
	// cú pháp the `channel <-`. Ở đây, chúng ta send
	// `"ping"` đến `messages` channel chúng ta tạo ra ở
	// trên, từ một goroutine mới.
	go func() { messages <- "ping" }()

	// Cú pháp `<-channel` _nhận_ một giá trị từ channel.
	// Ở dây, chúng ta sẽ nhận được gói tin `"ping"` mà
	// chúng ta đã gửi trước đó và in chúng ra.
	msg := <-messages
	fmt.Println(msg)
}
```
```sh
# Khi chúng ra chạy chương trình gói tin `"ping"`
# được truyền thành công từ một goroutine đến một
# goroutine khác thông qua channel của chúng ta.

$ go run channels.go 
ping

# Mặc định gửi và nhận bị chặn cho đến khi cả 
# channel gửi và nhận sẵn sàng. Thuộc tính này cho phép
# chúng ta đợi gói tin `"ping"` ở cuối chương trình mà 
# không phải sử dụng bất kì sự đồng bộ hóa nào khác. 
```
