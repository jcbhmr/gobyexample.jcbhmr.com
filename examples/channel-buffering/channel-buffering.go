// По умолчанию каналы _не буферизованы_, это означает,
// что они будут принимать отправления (`chan <-`), только
// если есть соответствующий прием (`<- chan`), готовый
// принять отправленное значение. _Буферизованные каналы_
// принимают ограниченное количество значений без
// соответствующего приемника для этих значений.

package main

import "fmt"

func main() {

	// Здесь мы `создаем` канал строк с буфером до 2
	// значений.
	messages := make(chan string, 2)

	// Т.к. этот канал буферизирован, мы можем послать
	// значения в канал без соответствующего одновременного
	// получения.
	messages <- "buffered"
	messages <- "channel"

	// Позже мы можем получить эти значения как обычно.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
