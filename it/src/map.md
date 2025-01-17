# Map
```go
// Le _Map_ sono la struttura built-in di Go per gli [Array associativi](https://it.wikipedia.org/wiki/Array_associativo)
// (in altri linguaggi si possono trovare strutture simili sotto
// il nome di _hash table_ o _dizionari_).

package main

import "fmt"

func main() {

    // Per creare una nuova map vuota, utilizza
    // la funzione built-in `make`:
    // `make(map[tipo-chiave]tipo-valore)`.
    m := make(map[string]int)

    // Puoi impostare i valori della map utilizzando la sintassi
    // tipica `nomemap[chiave] = valore`
    m["k1"] = 7
    m["k2"] = 13

    // Passare la map ad una funzione di stampa (tipo `Println`)
    // mostrerà tutte le coppie chiave-valore della map
    fmt.Println("map:", m)

    // Puoi ottenere il valore di una chiave con `nomemap[chiave]`.
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // La funzione built-in `len` restituisce il numero di coppie
    // chiave-valore se la si invoca su una map
    fmt.Println("len:", len(m))

    // La funzione built-in `delete` rimuove le coppie
    // chiave-valore dalla map
    delete(m, "k2")
    fmt.Println("map:", m)

    // Quando si accede ad una map è possibile controllare
    // il secondo valore restituito opzionale che indica
    // la presenza o meno di una chiave all'interno di una map.
    // Questo parametro può essere utilizzato per discernere
    // il caso in cui una chiave non è presente dal caso in cui
    // una chiave ha assegnato lo zero-value (ad esempio `0` o `""`).
    // In questo caso non abbiamo nemmeno bisogno del valore
    // associato alla chiave, per cui scartiamo il primo parametro
    // utilizzando l'identificatore _blank_ `_`
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // È anche possibile dichiarare ed inizializzare una
    // nuova map con la sintassi seguente
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
```
```sh
# Nota che le map vengono mostrate nel formato
# `map[k:v k:v]` se vengono stampate con `fmt.Println`.
$ go run maps.go 
map: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]
```
