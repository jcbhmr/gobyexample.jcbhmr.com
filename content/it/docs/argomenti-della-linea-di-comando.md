# Argomenti della Linea di Comando
```go
// Gli _argomenti della linea di comando_ sono un metodo
// comune per passare parametri ai programmi.
// Ad esempio, `go run hello.go` usa `run` e `hello.go`
// come parametri al programma `go`.

package main

import "os"
import "fmt"

func main() {

    // `os.Args` fornisce accesso dirretto agli argomenti
    // della linea di comando. Nota che il primo valore
    // in questa slice è il percorso del programma, e
    // `os.Args[1:]` contiene i veri argomenti del
    // programma.
    argsConProg := os.Args
    argsSenzaProg := os.Args[1:]

    // Puoi ottenere gli argomenti individualmente con
    // l'indexing.
    arg := os.Args[3]

    fmt.Println(argsConProg)
    fmt.Println(argsSenzaProg)
    fmt.Println(arg)
}
```
```sh
# To experiment with command-line arguments it's best to
# build a binary with `go build` first.
# Per sperimentare con gli argomenti della linea di
# comando, è meglio creare un file binario con `go build`
# prima.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Nel prossimo esempio vedremo come analizzare gli
# argomenti in un modo più avanzato con le flag.
```
