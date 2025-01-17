# Array
```go
// In Go, un _array_ è una sequenza numerata di elementi con una
// lunghezza specifica

package main

import "fmt"

func main() {

    // Qui creeremo un array chiamato `a` che conterrà esattamente
    // 5 elementi di tipo `int`. Il tipo degli elementi e la lunghezza
    // sono entrambi parti integranti del tipo dell'array. Per default
    // gli array vengono inizializzati allo _zero value_, che per gli
    // elementi di tipo `int` significa essere inizializzati a `0`.
    var a [5]int
    fmt.Println("emp:", a)

    // Possiamo assegnare un valore ad uno specifico indice utilizzando
    // la classica sintassi:
    // `array[indice] = valore` , e possiamo ottenere il valore con
    // `array[indice]`.
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // La funzione builtin `len` restituisce la lunghezza dell'array.
    fmt.Println("len:", len(a))

    // Utilizza questa sintassi per dichiarare ed inizializzare
    // un array nella stessa linea.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Gli array sono di base monodimensionali, ma possono essere composti
    // per costruire strutture di dati multidimensionali
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```
```sh
# Nota che gli array vengono visualizzati nella forma
# [v1 v2 v3 ...] se stampati con la funzione
# `fmt.Println`.
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# Noterai che gli _slice_ vengono utilizzati molto più
# spesso degli array in Go. Vedremo gli slices nel
# prossimo esempio.
```
