# Діапазон|Range
```go
// Ключове слово _range_ (діапазон) ітерує по елементах в деяких
// структурах даних доступних в Go. Давайте подивимось, як
// використовувати `range` з вже знайомими нам структурами даних.

package main

import "fmt"

func main() {

	// Ось - приклад сумування чисел [зрізу](./slices) (
	// з масивами це працює аналогічним чином).
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` на масивах та зрізах видає і індекс
	// і значення для кожного з входжень. У наведеному вище прикладі,
	// нам не потрібен був індекс, отож ми ігнорували його
	// за допомогою пустого ідентифікатора `_`. Але,
	// інколи, індекс все ж потрібен.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` застосований на мапу поверне ключ та значення
	// пари.
	kvs := map[string]string{"a": "абрикос", "b": "банан"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` також може просто ітерувати ключі на мапі.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` на рядку ітеруватиме через Unicode'ні коди.
	// Першим значенням буде початковий байт індекс руни (`rune`),
	// See [Strings and Runes](strings-and-runes) for more
	// а другим сама руна (`rune` - це тип значення в Go).
	for i, c := range "го" {
		fmt.Println(i, c)
	}
}
```
```sh
$ go run range.go
sum: 9
index: 1
a -> абрикос
b -> банан
key: a
key: b
0 103
1 111
```
