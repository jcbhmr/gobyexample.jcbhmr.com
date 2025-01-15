# Slice
```go
// Gli _Slice_ sono un data type fondamentale di Go,
// e rendono la gestione degli array più semplice e
// potente.

package main

import "fmt"

func main() {

    // A differenza degli array, gli slice vengono definiti
    // dando soltanto il tipo degli elementi che contengono
    // (non il numero di elementi). Per creare uno slice
    // vuoto con una lunghezza diversa da 0, usa la funzione
    // `make`. Di seguito creiamo uno slice di `string` di
    // lunghezza `3` (all'inizio zero-valued).
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // Possiamo impostare e prendere valori esattamente come
    // negli array.
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` restituisce, come ci si potrebbe aspettare, la
    // lunghezza dello slice.
    fmt.Println("len:", len(s))

    // Oltre a queste operazioni di base, gli slice ne
    // hanno molte altre che permettono loro di essere
    // più funzionali degli array. Una di queste è la
    // funzione `append`, che restituisce uno slice
    // contentente uno o più ulteriori valori.
    // Nota che abbiamo bisogno di accettare il valore
    // restituito da append, visto che potremmo ricevere
    // uno slice completamente nuovo.
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // Gli slice possono anche essere copiati con la
    // funzione `copy`. Di seguito creiamo uno slice
    // vuoto `c` della stessa lunghezza di s e copiamo
    // i valori di `s` in `c`.
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // Gli slice supportano un operatore "slice" che ha
    // la sintassi `variabileSlice[inizio:fine]`. Per
    // esempio, qui generiamo uno slice con gli
    // elementi `s[2]`, `s[3]` e `s[4]`.
    l := s[2:5]
    fmt.Println("sl1:", l)

    // Il seguente crea uno slice fino al quinto elemento
    // (escludendo il quinto).
    l = s[:5]
    fmt.Println("sl2:", l)

    // E il seguente lo crea degli elementi dopo il
    // secondo elemento (includendo il secondo).
    l = s[2:]
    fmt.Println("sl3:", l)

    // Possiamo, inoltre, dichiarare ed inizializzare
    // uno slice in una sola riga.
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // Gli slice possono essere composti in strutture di
    // dati a più dimensioni. Il numero degli elementi
    // degli slice all'interno può variare, a differenza
    // degli array multi-dimensionali.
    biDim := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        biDim[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            biDim[i][j] = i + j
        }
    }
    fmt.Println("2d: ", biDim)
}
```
```sh
# Nota che, anche se gli slice sono dei tipi diversi dagli
# array, anche essi possono essere stampati tramite
# `fmt.Println`.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Dai un'occhiata a questo [post](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)
# scritto dal team di Go per avere ulteriori dettagli sul
# design e la implementazione degli slice in Go.

# Ora che abbiamo visto array e slice, daremo un'occhiata
# all'altra struttura di dati basilare di Go: le _map_.
```
