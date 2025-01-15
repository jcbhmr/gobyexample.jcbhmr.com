# Panic
```go
// Un `panic` in genere significa che qualcosa è andato
// fin troppo storto. In genere lo usiamo per fallimenti
// immediati su errori che non dovrebbero succedere
// durante una normale esecuzione del programma, o per
// cui non è possibile effettuare procedure di recovery

package main

import "os"

func main() {

    // D'ora in poi utilizzeremo i `panic` su questo
    // sito per controllare errori inaspettati. Questo
    // è l'unico programma sul sito fatto apposta
    // perché risulti in un `panic`.
    panic("un problema")

    // Un uso comune del panic è l'interruzione del
    // programma quando ci troviamo davanti ad un errore
    // restituito da una funzione che non sappiamo come
    // gestire o non vogliamo gestire per niente. Di
    // seguito un esempio di `panic` quando riceviamo
    // un errore inaspettato dopo aver tentato di creare
    // un nuovo file.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
```
```sh
# Eseguire questo programma risulterà in un panic,
# il che significa che il messaggio di errore verrà
# stampato e anche gli stacktrace delle varie goroutine.
# Inoltre, il programma terminerà con uno status
# non-zero.
$ go run panic.go
panic: un problema

goroutine 1 [running]:
...
main.main()
	/.../panic.go:21 +0x65
exit status 2

# Nota che nonostante in molti linguaggi sia solito
# utilizzare le eccezioni per la gestione di tanti errori,
# in Go è idiomatico indicare gli errori in
# [un valore restituito a sé stante](errori) ove
# possibile.
```
