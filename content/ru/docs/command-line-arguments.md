# Аргументы командной строки (Command-Line Arguments)|Command-Line Arguments
```go
// [_Аргументы командной строки_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// являются распространенным способом параметризации
// выполнения программ. Например, `go run hello.go`
// использует аргументы `run` и `hello.go` для программы
// `go`.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `os.Args` предоставляет доступ к необработанным
	// аргументам командной строки. Обратите внимание,
	// что первое значение в этом срезе - это путь к
	// программе, а os.Args [1:] содержит аргументы
	// программы.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Вы можете получить отдельные аргументы с
	// обычной индексацией.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
```
```sh
# Чтобы поэкспериментировать с аргументами командной
# строки, лучше сначала создать двоичный файл с
# помощью `go build`.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Далее мы рассмотрим более сложную обработку командной
# строки с флагами.
```
