# Exit
```go
// Usa-se o `os.Exit` para, imediatamente, sair do programa
// com um determinado status.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` NÃO serão executados quando utilizado `os.Exit`,
	// então este `fmt.Println` nunca será chamado.
	defer fmt.Println("!")

	// Exit com status 3.
	os.Exit(3)
}

// Note que diferentemente de C, por exemplo, Go não
// utiliza um inteiro como valor de retorno de main,
// para indicar status de saída. Se a intenção for sair
// de um programa com status não-zero (sem sucesso),
// `os.Exit` deve ser usado.
```
```sh
#  Ao executar `exit.go` usando `go run`, a saída
# será capturada pelo Go e então impressa.
$ go run exit.go
exit status 3

# Ao construir e executar um binário, é possível
# visualizar o status no terminal.
$ go build exit.go
$ ./exit
$ echo $?
3

# Note que o `!` do programa nunca foi exibido.
```
