# Строковые функции (String Functions)|String Functions
```go
// Стандартная библиотека пакета `strings` предоставляет
// множество удобных функций для работы со строками. Вот
// некоторые из них.

package main

import (
	"fmt"
	s "strings"
)

// Создаем алиас для функции `fmt.Println`, т.к. далее мы будем
// часто вызывать эту функцию.
var p = fmt.Println

func main() {

	// Данные функции доступны в пакете `strings`. Обратите внимание,
	// что все эти функции из пакета, а не методы строковых объектов.
	// Это означает, что нам необходимо передать первым аргументом
	// функции, строку, над которой мы производим операцию. Вы можете
	// найти больше строковых функций в [`официальной документации
	// к пакету`](http://golang.org/pkg/strings/).
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// Примеры ниже не относятся к пакету `strings`, но о них
	// стоит упомянуть - это механизмы для получения длины
	// строки и получение символа по индексу.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Обратите внимание, что `len` и индексация выше работают на
// уровне байтов. Go использует строки в кодировке UTF-8, так
// что это часто полезно как есть. Если вы работаете с
// потенциально многобайтовыми символами, вам нужно использовать
// операции с кодировкой. Смотрите [строки, байты, руны и символы
// в Go](https://blog.golang.org/strings) для получения дополнительной
// информации.
```
```sh
$ go run string-functions.go
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:  5
Char: 101
```
