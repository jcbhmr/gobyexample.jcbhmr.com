# Epoch->时间戳
```go
// 一般程序会有获取 [Unix 时间](http://zh.wikipedia.org/wiki/UNIX%E6%97%B6%E9%97%B4)
// 的秒数，毫秒数，或者微秒数的需求。来看看如何用 Go 来实现。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 分别使用 `time.Now` 的 `Unix` 和 `UnixNano`，
	// 来获取从 Unix 纪元起，到现在经过的秒数和纳秒数。
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意 `UnixMillis` 是不存在的，所以要得到毫秒数的话，
	// 你需要手动的从纳秒转化一下。
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以将 Unix 纪元起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
```
```sh
$ go run epoch.go 
2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC

# 下面我们将看看另一个时间相关的任务：时间解析与格式化。
```
