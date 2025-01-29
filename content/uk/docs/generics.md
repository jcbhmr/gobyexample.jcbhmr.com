# Дженерики|Generics
```go
// Починаючи з версії 1.18, Go додає пітдтримку _дженериків_,
// також відомих як _параметри типу_.

package main

import "fmt"

// Як приклад узагальненої функції `MapKeys` приймає
// мапу будь-якого типу і повертає зріз її ключів.
// Ця функція має два параметри  `K` та `V`;
// `K` має `comparable` (порівняльне) _обмеження_, що
// означає що ми можемо порівнівати значення цього типу
// між собою за допомогою операторів `==` та `!=`. Це
// обмеження ключів у мапах Go. `V` має `any`
// обмеження, що означає - що тип не обмежений
// абсолютно (`any` є аліасом до `interface{}`).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Як приклад джерерика, `List` це зв'язаний список з
// будь-яким типом елементу.
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Ми також можемо визначати методи на узагальненних типах
// так само як ми це робили на звичайних, але ми мусимо слідкувати
// за параметрами, Тип `List[T]`, не `List`.
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	m := map[int]string{1: "2", 2: "4", 4: "8"}

	// Під час виклику узагальненої функції, ми можемо покластись
	// на _успадковування типу_. Зауважимо що нам не потрібно чітко
	// вказувати типи для `K` та `V` під час виклику `MapKeys` -
	// компілятор призначе їх автоматично.
	fmt.Println("keys:", MapKeys(m))

	// ...  але ми можемо визначати їх явно.
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
```
```sh
$ go run generics.go
keys: [4 1 2]
list: [10 13 23]

# Зауважимо: порядок ітерації по ключах карти не є
# визначеним у Go; отож різні виклики дадуть
# різні результати.
```
