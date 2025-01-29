# Rate Limiting
```go
// [_Rate limiting_](https://en.wikipedia.org/wiki/Rate_limiting)
// ou limitador de frequência é um mecanismo importante para
// controlar a utilização de recursos e manter a qualidade
// do serviço. Go suporta o rate limiting com goroutines,
// canais e [tickers](tickers) de maneira bem elegante.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Primeiro será apresentado rate limiting básico.
	// Suponha que seja necessário limitar o tratamento
	// das requisições recebidas, que serão servidas em
	// um canal `requests`.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Este canal limitador, ou `limiter`, receberá um valor
	// a cada 200 milisegundos. Este é o regulador no esquema
	// de rate limiting.
	limiter := time.Tick(200 * time.Millisecond)

	// Ao bloquear um recebimento do canal limitador
	// antes de atender a cada solicitação,
	// será tratada 1 solicitação por vez,
	// a cada 200 milissegundos.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Também é possível permitir envio de lotes de pequena
	// quantidade ou `short bursts` (rajadas curtas) de requisições
	// no esquema de rate limiting enquanto se cumpre com a
	// limitação de frequência no geral.
	// Para fazer isto, é necessário utilizar buffer no canal
	// `limiter`. Este canal `burstyLimiter` permitirá envios
	// de lotes de até 3 requisições.
	burstyLimiter := make(chan time.Time, 3)

	// Preenchendo o canal com requisições para representar
	// os lotes, ou bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200 milisegundos, será adicionado um valor
	// ao canal `burstyLimiter`, até o limite de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora é simulado outras 5 requisições sendo recebidas.
	// As 3 primeiras serão beneficiadas pela capacidade do
	//`burstyLimiter`de aceitar requisições em lote .
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
# Ao executar o código é fácil de visualizar o primeiro
# envio de requisições tratadas uma vez a cada
# ~200 milisegundos, como esperado.
$ go run rate-limiting.go
request 1 2012-10-19 00:38:18.687438 +0000 UTC
request 2 2012-10-19 00:38:18.887471 +0000 UTC
request 3 2012-10-19 00:38:19.087238 +0000 UTC
request 4 2012-10-19 00:38:19.287338 +0000 UTC
request 5 2012-10-19 00:38:19.487331 +0000 UTC

# Para o segundo conjunto de requisições, são tratadas
# de imediato as 3 primeiras, por conta do limitador
# em lotes e, só então as 2 remanescentes com intervalo
# de ~200 milisegundos entre cada uma.
request 1 2012-10-19 00:38:20.487578 +0000 UTC
request 2 2012-10-19 00:38:20.487645 +0000 UTC
request 3 2012-10-19 00:38:20.487676 +0000 UTC
request 4 2012-10-19 00:38:20.687483 +0000 UTC
request 5 2012-10-19 00:38:20.887542 +0000 UTC
```
