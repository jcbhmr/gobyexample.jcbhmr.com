# Hello World|Hello World
```go
// Наша первая программа напечатает классическое сообщение "hello world"
// Полный код.
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
```sh
# Для запуска программы, добавьте код в файл
# `hello-world.go` и выполните `go run`.
$ go run hello-world.go
hello world

# Иногда необходимо собрать программу в бинарный
# файл. Мы можем сделать это с помощью команды
# `go build`.
$ go build hello-world.go
$ ls
hello-world	hello-world.go

# Мы можем выполнить бинарный файл напрямую.
$ ./hello-world
hello world

# Теперь когда мы можем запускать и собирать Go-программы,
# давайте узнаем больше об этом языке.
```
