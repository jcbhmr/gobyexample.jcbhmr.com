# Сортировка (Sorting)|Sorting
```go
// Пакет `sort` реализует сортировку для встроенных и
// пользовательских типов. Сначала рассмотрим сортировку
// встроенных типов.

package main

import (
	"fmt"
	"sort"
)

func main() {

	// Методы сортировки специфичны для встроенного типа;
	// Вот пример для строк. Обратите внимание, что
	// сортировка выполняется на месте, поэтому она
	// изменяет данный фрагмент и не возвращает новый.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Пример сортировки `int`'ов
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// Мы так же можем использовать `sort`, для
	// проверки, что срез был уже отсортирован.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
```
```sh
# После запуска наша программа выведет отсортированные
# строки и срез целых чисел и `true`, как результат
# выполнения `AreSorted`.
$ go run sorting.go
Strings: [a b c]
Ints:    [2 4 7]
Sorted:  true
```