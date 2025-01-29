---
weight: 23
---
# Goroutines
```go
// Une _goroutine_ est un thread d'exécution léger.

package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // Suppons que l'on ait un appel de fonction `f(s)`.
    // Normalement, on l'appellerait comme ceci, et son
    // execution serait synchrone.
    f("direct")

    // Pour appeler la fonction dans une goroutine, on
    // utilise `go f(s)`. Cette nouvelle goroutine va
    // s'exécuter de manière concurrent avec l'appellée.
    go f("goroutine")

    // On peut aussi démarrer une goroutine pour un appel
    // de fonction anonyme.
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // Nous deux appels de fonction s'éxécutent de manière
    // asynchrone dans des goroutines séparées, donc
    // l'exécution tombe ici. Ce `Scanln` demande que
    // l'on presse une touche pour sortir du programme.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
```
```sh
# Lorsque l'on lance ce programme, on voit la sortie 
# de la fonction bloquante en premier, puis les 
# exécutions croisées des deux goroutines. 
# Cet entrelacement reflète le fait que les deux
# goroutines sont exécutées de manière concurrente.
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

# Ensuite nous allons regarder un complément aux 
# goroutines : les channels
```
