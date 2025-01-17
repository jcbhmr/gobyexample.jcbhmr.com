# HTTP Сервер|HTTP Server
```go
// Написання HTTP-серверу не є складною задачою,
// за умови використання пакет `net/http`.
package main

import (
	"fmt"
	"net/http"
)

// Фундаментальним концептом у серверах `net/http` є обробники (handlers)
// Обробник - це об'єкт, що імплементує інтерфейс `http.Handler`.
// Зазвичай, для написання обробника використовується адаптер `http.HandlerFunc`
// з функціює відповідної сигнатури.
func hello(w http.ResponseWriter, req *http.Request) {

	// Функції що використовуються як обробники - приймають
	// `http.ResponseWriter` і вказівник на `http.Request` як аргументи.
	// `http.ResponseWriter` використовується для запису HTTP-відповіді
	// В цьому прикладі - відповіддю є рядок "hello\n".
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	// В даному обробнику проводиться зчитування HTTP-заголовків,
	// і передавання їх як відповідь клієнту.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Дані рядки реєструють маршрути, які потрібно обслужити серверу,
	// за допомогою функції `http.HandleFunc`. Ця функція
	// встановлює маршрутизатор по замовчуванню і обробник який буде виконано,
	// якщо буде зроблено запит по вказаних маршрутах.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Запускаємо HTTP-сервер. В якості порту, на якому очікуються з'єднання
	// встановлено ":8090". Другий параметр вказує на те, що потрібно використовувати
	// маршрутизатор по замовчуванню, який було заздалегідь встановлено.
	http.ListenAndServe(":8090", nil)
}
```
```sh
# Run the server in the background.
$ go run http-servers.go &

# Зробити запит до маршруту "hello"
$ curl localhost:8090/hello
hello

# Завершимо наш сервер
$ kill -9 $!
```
