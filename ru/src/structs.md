# Структуры (Structs)|Structs
```go
// _Структуры_ в Go - это коллекции полей определенных
// типов. Как правило, они используются для логической
// группировки данных.

package main

import "fmt"

// Структура `person` имеет два поля `name` и `age`.
type person struct {
	name string
	age  int
}

// Функция NewPerson создает новую струкутуру person с
// заданным именем.
func NewPerson(name string) *person {
	// Вы можете безопасно вернуть указатель на локальную
	// переменную, так как локальная переменная переживет
	// область действия функции.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// Так создается новая структура
	fmt.Println(person{"Bob", 20})

	// Вы можете задавать имена для корректного
	// присваивания значений при создании структуры
	fmt.Println(person{name: "Alice", age: 30})

	// Пропущенные поля будут нулевыми.
	fmt.Println(person{name: "Fred"})

	// Префикс `&` возвращает указатель на структуру.
	fmt.Println(&person{name: "Ann", age: 40})

	// Можно инкапсулировать создание новой структуры
	// в функцию
	fmt.Println(NewPerson("Jon"))

	// Доступ к полям структуры осуществляется через
	// точку.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Вы также можете использовать точки со
	// структурными указателями - указатели автоматически
	// разыменовываются.
	sp := &s
	fmt.Println(sp.age)

	// Структуры мутабельны.
	sp.age = 51
	fmt.Println(sp.age)
}
```
```sh
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
Sean
50
51
&{Jon 42}
```
