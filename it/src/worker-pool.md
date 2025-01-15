# Worker Pool
```go
// In questo esempio vedremo come implementare
// un _worker pool_ usando le goroutine e i channel.

package main

import "fmt"
import "time"

// Questo è il worker, sul quale eseguiremo i nostri
// task concorrenti. Questi worker riceveranno i task da
// eseguire sul channel `jobs` ed invieranno i risultati sul
// channel `results`. In questo esempio abbiamo inserito una
// `Sleep` da un secondo per simulare un task oneroso.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "sta eseguendo task", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // Per poter utilizzare il pool di worker dobbiamo
    // poter inviare i task e poter ricevere i risultati,
    // creiamo dunque due channel per questo
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Questo farà partire 3 worker, che saranno inizialmente
    // bloccati in quanto non hanno task da eseguire.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Inviamo 9 task da eseguire sul channel `job` ed invochiamo
    // la `close` su quel canale, in modo da indicare che non
    // ci sono altri task da eseguire.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Infine recuperiamo i risultati delle computazioni sul
    // channel `results`
    for a := 1; a <= 9; a++ {
        <-results
    }
}
```
```sh
# Il nostro programma in esecuzione mostra i 9 task
# che vengono eseguiti dai vari worker. Il programma
# termina in circa 3 secondi, anche se il totale dei
# task avrebbe richiesto 9 secondi. Questo perchè ci
# sono 3 worker che vengono eseguiti in parallelo
$ time go run worker-pools.go 
worker 1 sta eseguendo task job 1
worker 2 sta eseguendo task job 2
worker 3 sta eseguendo task job 3
worker 1 sta eseguendo task job 4
worker 2 sta eseguendo task job 5
worker 3 sta eseguendo task job 6
worker 1 sta eseguendo task job 7
worker 2 sta eseguendo task job 8
worker 3 sta eseguendo task job 9

real	0m3.149s
```
