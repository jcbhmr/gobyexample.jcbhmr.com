# SHA256 Hashes
```go
// [_Giá trị băm SHA256 hashes_](https://en.wikipedia.org/wiki/SHA-2)
// thường được sử dụng để tính các định danh ngắn cho các nhị phân
// hoặc văn bản.  Ví dụ, chứng chỉ TLS/SSL sử dụng SHA256
// để tính chữ ký của chứng chỉ. Sau đây là cách tính giá trị băm
// SHA256 trong Go.

package main

//Go hiên thực một số hàm băm trong nhiều packages (gói)
// `crypto/*` khác nhau.
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// Ở đây, chúng ta bắt đầu với một đối tượng băm mới.
	h := sha256.New()

	// `Write` yêu cầu bytes. Nếu bạn có một chuỗi `s`,
	// sử dụng `[]byte(s)` để ép nó thành byte.
	h.Write([]byte(s))

	// Đây là kết quả băm đã hoàn thành dưới dạng byte slice.
	// Tham số của `Sum` có thể được sử dụng để thêm
	// vào một byte slice hiện có, thường không cần thiết.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
```
```sh
# Chạy chương trình tính toán giá trị băm và in nó
# dưới dạng hex có thể đọc được bởi con người.
$ go run sha256-hashes.go
sha256 this string
1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a...


# Bạn có thể tính toán các băm khác sử dụng một
# mẫu tương tự như mẫu được hiển thị ở trên. Ví dụ,
# để tính toán các giá trị băm SHA512, import 
# `crypto/sha512` và sử dụng `sha512.New()`.

# Lưu ý rằng nếu bạn cần các giá trị băm an toàn 
# bằng mật mã, bạn nên nghiên cứu kỹ
# [sức mạnh băm](https://en.wikipedia.org/wiki/Cryptographic_hash_function)!
```
