# Epoch|Epoch
```go
// Общим требованием в программах является получение
// количества секунд, миллисекунд или наносекунд в [Unixtime](http://en.wikipedia.org/wiki/Unix_time).
// Вот как это сделать в Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Используйте `time.Now` с `Unix` или `UnixNano`,
	// чтобы получить время, прошедшее с начала эпохи Unix в
	// секундах или наносекундах соответственно.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Обратите внимание, что `UnixMillis` не существует,
	// поэтому, чтобы получить миллисекунды с начала эпохи Unix,
	// вам нужно будет вручную делить наносекунды.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Вы также можете конвертировать целые секунды или наносекунды
	// Unixtime в соответствующее `время`.
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

# Далее мы рассмотрим еще одну задачу, связанную со
# временем: разбор и форматирование времени.
```