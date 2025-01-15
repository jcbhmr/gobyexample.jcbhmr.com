# Command-Line Arguments
```go
// [_Command-line arguments_](https://en.wikipedia.org/wiki/Command-line_interface#Arguments),
// ou argumentos de linha de comando,
// são uma forma comum de parametrizar execução de programas.
// Por exemplo `go run hello.go` usa como argumentos `run` e
// `hello.go` juntamente com o comando `go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` fornece acesso aos argumentos `raw`, ou crus
	// passados em linha de comando. Note que o primeiro valor
	// na slice é o caminho para o programa e `os.Args[1:]`
	// possui os argumentos de fato.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// É possível acessar argumentos individualmente com
	// indexação normal.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
```
```sh
# Para testar os argumentos de linha de comando,
# é melhor criar um binário com o comando `go build`
# antes.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Em seguida será apresentada a utilização
# de flags com linha de comando.
```
