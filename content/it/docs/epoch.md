# Epoch
```go
// Spesso i nostri programmi potrebbero aver bisogno di
// ottenere il numero di secondi, millisecondi o
// nanosecondi dallo <a href="https://it.wikipedia.org/wiki/Tempo_(Unix)">Unix Epoch</a>.
// Ecco come si può fare in Go.

package main

import "fmt"
import "time"

func main() {

    // Usa `time.Now` con `Unix` o `UnixNano` per
    // ottenere il tempo passato dallo Unix Epoch
    // rispettivamente in secondi o millisecondi.
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // Nota che non c'è nessun `UnixMillis`, quindi per
    // ottenere i millisecondi devi dividere manualmente
    // i nanosecondi.
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // Puoi anche convertire un int con il numero di
    // secondi o di millisecondi ad un `time.Time`.
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
```
```sh
$ go run epoch.go 
2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC

# Adesso daremo un'occhiata ad un altro task
# relativo al tempo: il parsing e la formattazione del
# tempo.
```
