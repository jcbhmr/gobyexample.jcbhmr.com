---
weight: 50
---
# Temps UNIX
```go
// Un besoin courant des programmes, c'est d'obtenir le nombre de secondes, millisecondes ou nanosecondes depuis l'[heure Unix](https://fr.wikipedia.org/wiki/Heure_Unix)
// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// [Unix epoch](http://en.wikipedia.org/wiki/Unix_time).
// Voici comment le faire en Go.

package main

import "fmt"
import "time"

func main() {

    // On utilise `time.Now` avec `Unix` ou `UnixNano` pour obtenir le temps écoulé depuis l'heure Unix en secondes ou en nanosecondes.
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // A noter qu'il n'y a pas de `UnixMillis`, donc pour obtenir des millisecondes il faut faire la division soit même à partir des nanosecondes.
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // On peut aussi convertir des secondes ou des nanosecondes depuis l'heure Unix.
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

# Ensuite, nous verrons une autre tâche liée au temps :
# parser et formatter des dates.
```
