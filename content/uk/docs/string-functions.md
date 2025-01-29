# Рядкові Функції|String Functions
```go
// Пакет стандартної бібліотеки - `strings` надає чимало корисних
// функцій, для роботи з рядками. Щоб показати можливості
// пакету `strings` приведемо кілька прикладів. Крім того -
// більше інформації можна отримати з документації на сторінці
// пакету [`strings`](http://golang.org/pkg/strings/).

package main

import (
	"fmt"
	s "strings"
)

// Скоротимо `fmt.Println` до `p` - оскільки використовувати
// ми його будемо багато, і так буде більш лаконічно.
var p = fmt.Println

func main() {

	// Ось приклади функцій, що доступні в `strings`.
	// Оскільки це функції пакету, а не власні методи
	// об'єкту рядка, нам необхідно передавати рядок першим
	// аргументом.
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
}
```
```sh
$ go run string-functions.go
ontains:   true
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
```
