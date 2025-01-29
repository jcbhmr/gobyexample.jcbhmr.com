# Маятники (Tickers)|Tickers
```go
// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Маятники використовують канал до якого надсилають значення,
	// це дуже схоже на те, як працюють таймери. Тут ми використаємо
	// `range` для каналу маятника і відразу виводимо значення
	// яке прибуває з каналу (раз на пів секунди).
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				f := "15:04:05.999999999"
				fmt.Println("Tick at", t.Format(f))
			}
		}
	}()

	// Після того як таймер зупинено (зауважте аналогію з зупинкою таймера),
	// з його каналу перестають надходити повідомлення.
	// Зупинимо наш таймер через 1600 мілісекунд
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```
```sh
# Ми бачимо 3 коливання після запуску програми та
# повідомлення про зупинку маятника.
$ go run tickers.go
Tick at 15:40:04.653271
Tick at 15:40:05.15445
Tick at 15:40:05.654003
Ticker stopped
```
