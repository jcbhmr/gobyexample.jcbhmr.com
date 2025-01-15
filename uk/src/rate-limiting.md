# Обмеження Частоти Запитів|Rate Limiting
```go
// [_Обмеження Частоти Запитів_](http://en.wikipedia.org/wiki/Rate_limiting)
// важливий механізм контролю за ресурсами, його використовують
// для підтримки необхідного рівня роботи (і обмеження навантаження
// на сервера з сторонніх джерел). Реалізація обмежувача в Go є можливою
// за допомоги [горутин](goroutines), [каналів](channels) та [маятників](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Розглянемо одну з реалізацій обмежувача частоти запитів.
	// Припустимо - нам необхідно обмежити обробку
	// вхідних запитів, отож, ми будемо працювати з ними (запитами)
	// через одноіменний канал (`requests`).
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Канал `limiter` (або "обмежувач") буде приймати
	// значення кожні 200 мілісекунд - це так званий "маятник",
	// або "регулятор" нашої системи обмеження.
	limiter := time.Tick(200 * time.Millisecond)

	// Тут відбувається блокування на отримання з каналу
	// `limiter` перед обробкою кожного запиту, ми чекатимемо
	// повідомлення кожні 200 мс і обробляємо по одному запиту
	// за цей період.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Можливо, ми захочемо дозволяти невеликі "навали" запитів,
	// в цілому притримуючись нашої стандартної схеми обмежень.
	// Досягнути цього ми можемо створивши буфер
	// в нашому каналі "обмеження". Канал `burstyLimiter`
	// дозволятиме напливи, аж, трьох подій.
	burstyLimiter := make(chan time.Time, 3)

	// Заповнено канал значеннями, для імітування "напливу".
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Кожні 200 мілісекунд ми спробуємо додати значення в
	// `burstyLimiter`, аж до його ліміту в 3 події. Ми, так
	// би кажучи - знову "наповнюємо пустий келих".
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Ми симулюємо 5 вхідних подій, 3 з яких
	// - не будуть обмежені у виконанні, а все тому,
	// що вони отримали перевагу від буферизованого
	// каналу `burstyLimiter`, який блокує доступ
	// для 4-того та 5-того запитів, аж до моменту коли
	// сам канал "обмежувач" не буде знову наповнено.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```
```sh
# Запускаючи нашу програму - ми бачимо першу
# партію запитів, які обробляються кожні 200
# мілісекунд.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Друга ж партія, представляє собою
# три запити що не обмежуються (це "навала"
# - дозволена кількість без обмеження) та
# ще 2 - додаткових, що обмежені 200
# мілісекундними інтервалами.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```