# Tickers
```go
// [Timers](timers), como já dito, servem para executar
// alguma coisa, uma vez, em algum momento futuro.
// _Tickers_ por sua vez, servem para executar alguma
// coisa, repetidamente, em intervalos regulares.
// Aqui está um exemplo de ticker que roda periodicamente
// até ser interrompido.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Tickers usam um mecanismo parecido com timers: um
	// canal para o qual valores são enviados. Aqui será
	// utilizado o recurso `select` para aguardar os
	// valores que chegarão a cada 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers também podem ser interrompidos como timers.
	// Uma vez que um ticker é interrompido, não receberá
	// mais nenhum valor no seu canal.
	// Este será parado depois de 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```
```sh
# Ao rodar este programa, o ticker deve executar 
# 3 vezes antes de ser interrompido.
$ go run tickers.go
Tick at 2012-09-23 11:29:56.487625 -0700 PDT
Tick at 2012-09-23 11:29:56.988063 -0700 PDT
Tick at 2012-09-23 11:29:57.488076 -0700 PDT
Ticker stopped
```