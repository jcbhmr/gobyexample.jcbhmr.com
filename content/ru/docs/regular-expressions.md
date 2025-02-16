# Регулярные выражения (Regular Expressions)|Regular Expressions
```go
// Go предлагает встроенную поддержку [регулярных выражений](http://en.wikipedia.org/wiki/Regular_expression).
// Вот несколько примеров, связанных с регулярными
// выражениями в Go.

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Проверяем, соответствует ли шаблон строке
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Выше мы использовали строковый шаблон напрямую,
	// но для других задач с регулярными выражениями,
	// необходимо `скомпилировать` оптимизированную
	// структуру `Regexp`.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Множество методов доступны для этой структуры.
	// Вот тест на совпадение, который мы видели ранее.
	fmt.Println(r.MatchString("peach"))

	// Этот метод находит соответствие для регулярного
	// выражения.
	fmt.Println(r.FindString("peach punch"))

	// Этот метод также находит первое совпадение, но
	// возвращает начальный и конечный индексы совпадения
	// вместо текста.
	fmt.Println(r.FindStringIndex("peach punch"))

	// Варианты `Submatch` включают в себя информацию
	// как о совпадениях с полным шаблоном, так и о
	// совпадениях с частями шаблона. Например, эта
	// конструкция вернет информацию как
	// для `p([a-z]+)ch`, так и для `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Точно так же это возвратит информацию об индексах
	// совпадений и подсовпадений.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Метод `All` применяется ко всем совпадениям на входе,
	// а не только к первому. Например, чтобы найти все
	// совпадения для регулярного выражения.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Этот метод `All` доступен и для других функций,
	// которые мы видели выше.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Указание неотрицательного целого числа в качестве
	// второго аргумента для этих функций ограничит
	// количество совпадений.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// В наших примерах выше были строковые аргументы и
	// использовались такие имена, как `MatchString`. Мы
	// также можем предоставить `[]byte` аргументы и удалить
	// `String` из имени функции.
	fmt.Println(r.Match([]byte("peach")))

	// При создании констант с регулярными выражениями
	// вы можете использовать `MustCompile`, как аналог
	// `Compile`. Обычный `Compile` не будет работать
	// для констант, потому что он возвращает 2 значения.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Пакет `regexp` также можно использовать для
	// замены подмножеств строк другими значениями.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Вариант с `Func` позволяет вам преобразовывать
	// сопоставленный текст с заданной функцией.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
```
```sh
$ go run regular-expressions.go 
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH

# Для получения полной ссылки на регулярные выражения Go
# проверьте документацию пакета [`regexp`](http://golang.org/pkg/regexp/).
```
