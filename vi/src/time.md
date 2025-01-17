# Time
```go
// Go hỗ trợ rất nhiều về tính toán thời gian và khoảng thời gian;
// đây là một số ví dụ.
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Ta sẽ bắt đầu bằng việc tìm thời gian hiện tại.
	now := time.Now()
	p(now)

	// Ta có thể tạo một struct `time` bằng cách cung cấp năm,
	// tháng, ngày, v.v. Thời gian luôn được liên kết với một
	// `Location`, ví dụ như múi giờ.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Ta có thể tách các thành phần khác nhau
	// của thời gian ra như mong muốn.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Các ngày trong tuần từ thứ 2 đến Chủ Nhật có thể được tách ra.
	p(then.Weekday())

	// Các phương thức sau so sánh hai thời điểm, kiểm tra
	// xem thời điểm đầu tiên xảy ra trước, sau, hoặc cùng
	// thời điểm với thời điểm thứ hai.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Phương thức `Sub` trả về một `Duration` biểu diễn
	// khoảng thời gian giữa hai thời điểm.
	diff := now.Sub(then)
	p(diff)

	// Ta có thể tính toán độ dài của khoảng thời gian
	// theo các đơn vị khác nhau.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Ta có thể sử dụng `Add` để tăng thời gian lên một khoảng
	// thời gian cho trước, hoặc sử dụng cùng với `-` để
	// giảm thời gian đi một khoảng thời gian cho trước.
	p(then.Add(diff))
	p(then.Add(-diff))
}
```
```sh
$ go run time.go
2012-10-31 15:50:13.793654 +0000 UTC
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
25891h15m15.142266763s
25891.25420618521
1.5534752523711128e+06
9.320851514226677e+07
93208515142266763
2012-10-31 15:50:13.793654 +0000 UTC
2006-12-05 01:19:43.509120474 +0000 UTC

# Tiếp theo chúng ta sẽ xem các cách xử lý về 
# thời gian liên quan đến Unix epoch.
```
