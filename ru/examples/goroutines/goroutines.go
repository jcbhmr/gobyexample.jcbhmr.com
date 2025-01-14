// _Горутины_ - это легковесный тред.

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Предположим, у нас есть вызов функции `f(s)`. Вот
	// как мы бы вызвали её обычным способом, запустив
	// синхронно.
	f("direct")

	// Чтобы вызвать эту функцию в горутине, используйте
	// `go f(s)`. Эта новая горутина будет выполняться
	// одновременно с вызывающей фукнцией.
	go f("goroutine")

	// Вы так же можете вызывать анонимные функции в
	// горутнах.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Наши две функции теперь выполняются асинхронно в
	// отдельных горутинах. Дождитесь их окончания (для
	// более надежного подхода используйте [WaitGroup](waitgroups)).
	time.Sleep(time.Second)
	fmt.Println("done")
}