# Goroutines
```go
// Uma _goroutine_ é uma leve _thread_ de execução.

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

	// Supondo que se pretenda chamar a função `f(s)`.
	// Aqui está como seria feito para rodar a função
	// sincronamente.
	f("direct")

	// Para invocar esta função de maneira assíncrona,
	// em uma goroutine, usa-se `go f(s)`. Esta Goroutine
	// irá executar de maneira concorrente.
	go f("goroutine")

	// Também é possível iniciar uma goroutine a partir
	// de uma chamada de função anônima.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// As duas chamadas de funções estão rodando assincronamente em
	// goroutines separadas. Para uma apresentação mais detalhada,
	// veja [WaitGroups](waitgroups).
	time.Sleep(time.Second)
	fmt.Println("done")
}
```
```sh
# Ao rodar este programa, primeiro é visualizada a saída
# das chamadas bloqueantes, síncronas, e só então a saída 
# das duas goroutines. A saída de goroutines podem ser
# exibidas em ordem diversa da disposta no código,
# justamente porque goroutines são executadas 
# concorrentemente pelo runtime de Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
done

# Em seguida veremos um complemento às goroutines
# em programas Go que utilizam concorrência: canais.
```
