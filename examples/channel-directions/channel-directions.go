// Використовуючи [канал](channels) як аргумент функції,
// ми можемо вказати його призначення: буде він використаний
// для відправки, чи для отримання значень. Це дуже корисно з
// точки зору додадкової безпеки типів значень

package main

import "fmt"

// Функція `ping` прийме канал, лише, для надсилання значень.
// Ви отримаєте помилку компіляції, якщо спробуєте отримати
// значення з цього каналу.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Функція `pong` прийме перший канал для надсилання
// (`pings`), a другий для відправки (`pongs`) значень.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
