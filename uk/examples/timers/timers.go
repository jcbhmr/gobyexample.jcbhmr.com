// We often want to execute Go code at some point in the
// future, or repeatedly at some interval. Go's built-in
// _timer_ and _ticker_ features make both of these tasks
// easy. We'll look first at timers and then
// at [tickers](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Хронометр - це подія що відбудеться в майбутньому. Ви вказуєте
	// хронометру скільки часу ви хочете зачекати до неї, а він забезпечує
	// вам канал, по якому він передасть повідомлення, що час сплив.
	// Наприклад, наш хронометр розрахований на дві секунди.
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` блокується каналом `C` допоки не буде
	// надіслано повідомлення що час минув.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// Якщо вам треба зачекати, ви можете скористатись просто
	// `time.Sleep`.  Крім того, не забувайте, хронометр
	// можна зупинити в будь-який момент.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Даємо `timer2` достатньо часу для запуску (ніби
	// він колинебудь запуститься) щоб продемонструвати що
	// він вже зупиненний.
	time.Sleep(2 * time.Second)
}
