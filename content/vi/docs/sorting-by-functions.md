# Sorting by Functions
```go
// Thi thoảng chúng ta sẽ muốn sắp xếp một collection
// theo một thứ gì đó ngoài thứ tự tự nhiên. Ví dụ, giả
// sử chúng ta muốn sắp xếp các chuỗi theo độ dài thay
// vì theo thứ tự bảng chữ cái. Đây là một ví dụ về
// các loại sắp xếp tùy chỉnh trong Go.

package main

import (
	"fmt"
	"sort"
)

// Để sắp xếp theo hàm tùy chỉnh trong Go, chúng ta
// cần type tương ứng. Ở đây, chúng ta đã tạo một kiểu
// `byLength` chỉ là một alias cho kiểu builtin
// `[string]`.
type byLength []string

// Chúng ta triển khai `sort.Interface` - `Len`, `Less`
// và `Swap` - Trên các loại chúng ta có thể sử dụng
// gói `sort` hàm `Sort` chung. `Len` và `Swap` sẽ
// thường sẽ giống nhau giữa các type và `Less` sẽ
// giữ logic sắp xếp tùy chỉnh thực tế. Trong trường hợp
// chúng ta muốn sắp xếp theo thứ tự độ dài của chuỗi,
// vậy chúng ta dùng `len(s[i])` và `len(s[j])` ở đây.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Với tất cả những điều trong này, giờ chúng ta có thể
// thực hiện sắp xếp tùy chỉnh bằng cách chuyển đổi
// slice `fruits` to `byLength`, và sau đó sử dụng
// `sort.Sort` trên kiểu slice.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
```
```sh
# Chạy chương trình bày ra một danh sách đã được sắp xếp
# theo độ dài chuỗi, như mong muốn.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# Bằng cách làm theo cùng một mô hình tạo tùy chỉnh này
# gõ, triển khai ba phương thức `Interface` trên đó
# gõ và sau đó gọi sort.Sort trên một collection của loại
# tùy chỉnh đó, chúng ta có thể sắp xếp Go slice tùy ý
# chức năng.
```
