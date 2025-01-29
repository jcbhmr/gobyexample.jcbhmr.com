# Відновлення|Recover
```go
// Функція _recover_ (або відновлення) використовується лише
// в парі з _defer_, її призначення перехоплювати _panic_
// та відновлювати потік виконання програми.

// An example of where this can be useful: a server
// wouldn't want to crash if one of the client connections
// exhibits a critical error. Instead, the server would
// want to close that connection and continue serving
// other clients. In fact, this is what Go's `net/http`
// does by default for HTTP servers.

package main

import "fmt"

// Основна мета цієї функції - викликати `паніку`.
func mayPanic() {
	panic("Паніка без причини")
}

func main() {

	// Перехоплення _panic_'и та _recover_ відбувається
	// за допомогою `відкладеного` виклику, який, як ви вже
	// знаєте, відбувається по завершенню виконання функції.
	defer func() {
		if r := recover(); r != nil {
			// Значення поверненне recover i є помилка з якою програма
			// панікує.
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// Цей код незапуститься, оскільки `mayPanic` панікує.
	// Запуск `main` зупиняється на моменті panics і відновлюється у
	// відкладеному замиканні.
	fmt.Println("After mayPanic()")
}
```
```sh
$ go run recover.go
Recovered. Error:
  Паніка без причини
```
