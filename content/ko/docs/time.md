# Time->시간
```go
// Go는 시간(times)과 기간(durations)에 대한 광범위한 지원을 제공합니다.
//  여기에 몇 가지 예시가 있습니다.

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// 현재 시간을 구하는것부터 시작합시다.
	now := time.Now()
	p(now)

	// 연, 월, 일등을 전달하여 `time` 구조체를 생성할 수 있습니다.
	//  시간은 언제나 `Location` 즉, 타임존과 관련이 있습니다.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 시간값으로부터 다양한 컴포넌트를 가져올 수 있습니다.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 월요일부터 일요일까지의 '주일'도 사용 가능합니다.
	p(then.Weekday())

	// 다음 메서드들은 두 시간을 비교하는데, 첫번째값이 두번째값보다 이전인지, 이후인지 또는 같은지를 검사합니다.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// `Sub` 메서드는 두 시간 사이의 간격을 나타내는 `Duration`을 반환합니다.
	diff := now.Sub(then)
	p(diff)

	// 다양한 단위로 기간의 길이를 계산할 수 있습니다.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// `Add`를 사용해 주어진 기간만큼 시간을 앞당길 수 있으며, `-`를 사용하면 뒤로 이동할 수 있습니다.
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

# 다음으로 유닉스 Epoch과 관련된 시간과 관련된 아이디어를 살펴보겠습니다.
```
