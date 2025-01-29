# Panic
```go
// Um `panic` tipicamente significa que acontenceu algum
// erro inesperado. Geralmente é utilizado para falhar
// rapidamente em erros que não deviam ocorrer durante
// uma operação normal, ou que não está sendo tratado
// de maneira _graciosa_ ou apropriada.

package main

import "os"

func main() {

	// O panic será utilizado por todo o site para checar por
	// erros inesperados. Este é o único código em todo o site
	// especificamente desenhado para causar um panic.
	panic("uhhh... Houston, we have a problem.")

	// Um uso comum do panic é para abortar a execução
	// de determinado código se uma função retorna um
	// erro que não é tratado. Aqui está um exemplo
	// simulado de `panic` ao receber um erro inesperado
	// ao se criar um novo arquivo.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
```sh
# Ao executar este código ocorrerá um panic e será
# exibido no console uma mensagem de erro, um 
# "rastro" de execução da goroutine, e saída, 
# ou exit, com código diferente de zero.

# Quando o primeiro panic na função main dispara, 
# o código é interrompido, sem continuar até o 
# final da execução. Para visualizar o que ocorre
# no segundo panic deste exemplo, é necessário
# comentar o primeiro panic.
$ go run panic.go
panic: uhhh... Houston, we have a problem.

goroutine 1 [running]:
main.main()
	/.../panic.go:16 +0x27
exit status 2


# Note que, diferentemente de outras linguagens que 
# usam exceções para tratar erros, em Go é idiomático
# usar retornos indicativos de erro sempre que possível.
```
