# Вбудовування структур|Struct Embedding
```go
// Go підтримую _вбудовування_ структур та інтерфейсів
// як пирклад композиції типів. Не плутайте будь ласка з
// директивою [`//go:embed`](embed-directive) що була
// введена у 1.16+ для вбудовування у бінарник файлів.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` _вбудовує_ `base`. Ось так виглядає
// вбудовуванння без імені.
type container struct {
	base
	str string
}

func main() {
	// Коли ми створємо структури з двох літералів, ми маємо яквно
	// ініціалізувати убудування. Тут ми наприклад вбудований тип
	// служить як імя поля.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Ми маємо доступ до поля base прямо через головний тип `co`,
	// напряклад `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// альтернативно ми можемо і прописати повний шлях
	// використовуючи тип вбудованого обєкта.
	fmt.Println("also num:", co.base.num)

	// І оскільки `container` вбудовує `base`, методи останнього стають
	// також методами `container`. Напряклад ми викликаємо метод що було
	// вбудовано з base.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Вбудова структур з методами може бути використаня для
	// побудови імплементацій інтерфейсів у інші структури.
	// Тут ми бачимо `container` що імплементує інтерфейс `descriptor`
	// тому що він має вбудований `base`
	var d describer = co
	fmt.Println("describer:", d.describe())
}
```
```sh
$ go run struct-embedding.go
co={num: 1, str: some name}
also num: 1
describe: base with num=1
describer: base with num=1
```
