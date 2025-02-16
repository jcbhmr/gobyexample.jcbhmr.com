# Структури|Structs
```go
// Структури в Go - це колекції полів з визначеним типом.
// Надзвичайну корисність структур можна оцінити не тільки для
// групування даних, а й тому що вони служать основним рушієм мови
// що орієнтується на данні - на структури.

package main

import "fmt"

// Структура `person` має поля для імені (`name`) та віку (`age`).
type person struct {
	name string
	age  int
}

// `newPerson` створює нову персону з заданим ім'ям.
func newPerson(name string) *person {
	// Можна безпечно повертати вказівник на локальну змінну
	// і ця змінна переживе свій скоуп.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Дозволяється пропускати поля при ініціалізації структури.
	fmt.Println(person{"Bob", 20})

	// Дозволяється іменувати поля при ініціалізації структури.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// Префікс `&` поверне вказівник на структуру.
	fmt.Println(&person{name: "Ann", age: 40})

	// Приховувати процес створення стуктури іншою функцією - це
	// ідіоматичний підхід у Go.
	fmt.Println(newPerson("Jon"))

	// Доступ до полів надається через синтаксис крапки `.`
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Дозволяється використовувати крапки з вказівниками структур,
	// вказівники, в такому разі, автоматично розіменовуються.
	sp := &s
	fmt.Println(sp.age)

	// Дані у структурі можна змінювати (тобто вони `mutable`,
	// такі що дозволяється змінювати).
	sp.age = 51
	fmt.Println(sp.age)

	// Якщо структура використовувається лише короткостроково, ми
	// не даємо їй імені. Значення може містити анонімну структуру,
	// ця тезніка часто використовуєтся для
	// [табличних тестів](testing-and-benchmarking).
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
```
```sh
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
{Rex true}
```
