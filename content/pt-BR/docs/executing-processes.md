# Executing Processes
```go
// No exemplo anterior foi apresentada uma forma de
// [invocar processos externos](spawning-processes).
// Isto é feito quando é necessário um processo externo
// acessível a um programa Go sendo executado.
// Algumas vezes, no entanto, é necessário substituir
// completamente o processo Go corrente, por um outro
// (talvez até não-Go). Para fazer isto, será utilizada
// a implementação da clássica função
// <a href="https://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>.

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Para este exemplo, será executado `ls`. Go requer um
	// caminho absoluto para o binário que se pretende executar,
	// então será usado `exec.LookPath` para localizar (provavelmente será
	// `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` requer argumentos no formato de slice (e
	// não na forma de string). Será passado ao `ls`
	// alguns argumentos comuns. Note que o primeiro
	// argumento deve ser o nome do programa.
	args := []string{"ls", "-a", "-l", "-h"}

	// `Exec` também necessita um conjunto de [variáveis de ambiente](environment-variables)
	// para usar. Aqui é fornecido o ambiente atual
	// por completo.
	env := os.Environ()

	// Aqui está a chamada à `syscall.Exec`. Se esta
	// chamada for bem sucedida, a execução do processo
	// terminará e será substituído pelo processo
	// `/bin/ls -a -l -h`. Se houver algum erro, será
	// retornado um valor.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
```
```sh
# Ao rodar o programa, ele é substituído pelo `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Note que Go não oferece a função `fork`,
# clássica do Unix. Usualmente isto não é um
# problema, porque goroutines, invocação e
# executação de processos cobrem a maioria dos
# casos de uso de `fork`.
```
