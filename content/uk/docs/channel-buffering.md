# Буферизація каналiв|Channel Buffering
```go
// Відповідно до стандартних налаштувань, канали не _буферизуються_,
// тобто вони є такими - що тільки прийматимуть повідомлення (`chan <-`)
// якщо відповідний отримувач (`<- chan`) готовий прийняти його.
// _Буферизовані канали_ приймуть обмежену кількість значень,
// не чекаючи того щоб отримувач забрав повідомлення.

package main

import "fmt"

func main() {

	// Ось ми створюємо канал рядків, з буфером в 2 значення.
	messages := make(chan string, 2)

	// Оскільки канал використовує буфер, ми можемо надіслати до нього
	// оці значення, без відповідного читання з іншого кінця
	// каналу.
	messages <- "buffered"
	messages <- "channel"

	// І ми можемо отримати ці значення коли нам заманеться.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```
```sh
$ go run channel-buffering.go 
buffered
channel
```
