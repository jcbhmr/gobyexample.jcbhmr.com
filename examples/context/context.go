// У попередньому прикладі ми розглядали як налаштувати простий
// [HTTP сервер](http-servers). Скористаємось HTTP сервером для демонстрації `context.Context` для контрольованого завершення.
// `Context` проносить в собі часові обмеження, сигнали зупинки,
// та інші деталі в рамках окремого запиту через API та горутини.
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {

	// `context.Context` створено для кожного запиту
	// механікою пакету `net/http`, і доступно через
	// метод `Context()`.
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// Зачекаймо кілька секунд перед надсиланням відповіді клієнту.
	// Це може симулювати деяку роботу серверу, поки працюєте -
	// спостерігайте за `Done()`, який має доповісти про сигнал
	// відміни роботи і припринення як умога раніше.
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// Метод контексту `Err()` повертає помилку, що
		// посянює чому канал `Done()` було закрито.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	// Як і раніше, ми реєструємо обробник на шлях "/hello"
	// <nobr>та починаємо обслуговування.</nobr>
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
