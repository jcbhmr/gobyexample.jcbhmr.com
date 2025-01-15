# Goroutine
```go
// Una _goroutine_ è di fatto un [green thread](https://en.wikipedia.org/wiki/Green_threads) che permette di
// rendere concorrente l'esecuzione del nostro programma.

package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // Supponiamo di avere una chiamata di funzione `f(s)`.
    // Questo è il modo classico di invocarla, che la eseguirà
    // in modo sincrono.
    f("direct")

    // Per invocare questa funzione in una goroutine, utilizza
    // `go f(s)`. Questa nuova goroutine eseguirà in parallelo alla
    // goroutine dove è stata invocata.
    go f("goroutine")

    // È possibile far partire una goroutine anche con una funzione anonima.
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // Le nostre due chiamte di funzione stanno eseguendo in modo asincrono,
    // su due goroutine separate. La funzione `Scanln` richiede che si prema un tasto,
    // prima di far terminare l'esecuzione della goroutine principale
    // (che terminerà tutte le altre goroutine).
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
```
```sh
# Quando si esegue questo programma si può notare prima
# l'output delle chiamate bloccanti. Successivamente si
# nota come le due goroutine lanciate producano
# dell'output "disordinato". Questo disordine riflette
# l'esecuzione concorrente delle goroutine da parte del
# runtime di Go.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# Adesso siamo pronti per analizzare la struttura
# complementare alle goroutine in Go: i channel.
```
