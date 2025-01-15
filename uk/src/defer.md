# Відкладення|Defer
```go
// Конструкція `defer` або _відкладення_ використовують
// щоб впевнитись, що виклик функції буде виконано пізніше,
// по завершенню роботи програми, найчастіше - для
// звільнення залучених ресурсів. `defer` вживають там,
// де `ensure` та `finally` були б використані в
// інших мовах.

package main

import (
	"fmt"
	"os"
)

// Припустимо, нам потрібно створити файл, записати
// до нього якусь інформацію, і коли ми закінчимо
// роботу з ним - закрити його. Ми _відкладемо_ закриття
// файлу, за допомогою інструкції `defer`.
func main() {

	// Після створення об'єкту `File` ( результат роботи `createFile` ),
	// ми відкладаємо його закриття за допомогою `closeFile`.
	// Виклик `closeFile` буде виконано по закінченні
	// роботи функції `main`, але вже після того,
	// як `writeFile` завершить свою роботу.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// Важливо виконувати перевірку на помилки, підчас завершення роботи
	// з файлом, навіть у відкладеній функції.
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
```
```sh
# Запуск програми доводить - файл було
# закрито після того, як до нього було записано дані.
$ go run defer.go
creating
writing
closing

# Тепер нам відомі такі поняття як
# [паніка (`panic`)](panic) та відкладення
# (defer), і ми готові ознайомитись з концептом
# [`відновлення (recover)`](recover).
```
