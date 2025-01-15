# Ограничение скорости (Rate Limiting)|Rate Limiting
```go
// <em>[Ограничение скорости](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// является важным механизмом контроля использования ресурсов и
// поддержания качества обслуживания. Go элегантно поддерживает
// ограничение скорости с помощью горутин, каналов и [тикеров](tickers).

package main

import (
	"fmt"
	"time"
)

func main() {

	// Сначала мы рассмотрим базовое ограничение скорости.
	// Предположим, мы хотим ограничить нашу обработку
	// входящих запросов. Мы будем обслуживать эти запросы
	// с одноименного канала.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Канал `limiter` будет получать значение каждые
	// 200мс. Это то, что регулирует скорость в нашей
	// схеме.
	limiter := time.Tick(200 * time.Millisecond)

	// Блокируя прием от канала `limiter` перед
	// обслуживанием каждого запроса, мы ограничиваем себя
	// одним запросом каждые 200 миллисекунд.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Мы можем разрешить короткие всплески запросов в
	// нашей схеме ограничения скорости при сохранении
	// общего ограничения скорости. Мы можем сделать это
	// путем буферизации нашего канала ограничения. Этот
	// канал `burstyLimiter` будет позволять делать до
	// 3 событий.
	burstyLimiter := make(chan time.Time, 3)

	// Заполняем канал, чтобы предоставить возможность
	// ускорить.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Каждые 200мс мы будем пытаться добавлять новое
	// значение в `burstyLimiter`, до своего предела
	// в 3 значения.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Теперь смоделируем еще 5 входящих запросов. Первые
	// 3 из них получат выгоду от вместимости `burstyLimiter`.
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
# При запуске нашей программы мы видим, что первая
# партия запросов обрабатывается каждые ~200мс.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Для второго пула запросов мы обслуживаем первые
# 3 сразу из-за использования ограничения скорости,
# затем обслуживаем оставшиеся 2 с задержками ~200мс
# каждый.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```