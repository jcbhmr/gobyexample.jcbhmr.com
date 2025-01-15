# Panic
```go
// Một `panic` thương có nghĩa là một cái gì đó xảy ra
// không như mong đợi. Hầu hết chúng ta sử dụng nó để
// fail fast với các lỗi không nên xảy ra trong quá
// trình hoạt động bình thường, hoặc rằng chúng ta không
// được chuẩn bị để xử lý một cách tốt.

package main

import "os"

func main() {
	// Chúng ta sẽ dùng panic trên khắp site để kiểm tra
	// những lỗi không mong đợi. Đây là chương trình duy
	// nhất trên site được thiết kể để panic.
	panic("a problem")

	// Một trường hợp thông thường của panic là để hủy
	// bỏ nếu một hàm trả về một giá trị error mà chúng
	// ta không biết làm thế nào (hoặc muốn) để xử lý.
	// Đây là ví dụ của `panic` nếu chúng ta lấy
	// một lỗi xảy ra không mong đợi khi tạo file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
```sh
# Chạy chương trình sẽ gây ra panic, in một thông báo
# lỗi và dấu vết goroutine, và thoát với một trạng 
# thái non-zero.


# Khi panic đầu tiên trong `main` được bắn ra, chương
# trình thoát mà không đạt đến phần còn lại của code.
# Nếu bạn muốn để xem chương trình hãy thửu tạo một
# tệp tạm thời, comment lại cái panic đầu tiên.

$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
    /.../panic.go:12 +0x47
...
exit status 2

# Lưu ý rằng không giống như một số ngôn ngữ sử dụng
# exception để xử lý nhiều lỗi, trong Go nó là idiomatic 
# để sử dụng các giá trị trả về chỉ báo lỗi bất cứ
# khi nào có thể.
```
