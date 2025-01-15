# Регулярні Вирази|Regular Expressions
```go
// Як і інші мови програмування - Go має підтримку
// [регулярних виразів](https://uk.wikipedia.org/wiki/Регулярний_вираз)
// у стандартній бібліотеці. Ось, деякі приклади задач, що так чи
// інакше пов'язані з регулярними виразами (або `regexp` - як їх зазвичай називають).

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	// Тестуємо - чи наш шаблон співпадає з рядком.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Рядком вище - ми скористались рядковим шаблоном напряму,
	// але для інших задач, пов'язаних з регулярними виразами,
	// нам необхідно буде `Compile` (скомпілювати) рядковий
	// шаблон у `Regexp` структуру.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// На таких структура доступно чимало різних методів:
	// ось наприклад співставлення з рядком, що ми бачили в
	// першому прикладі.
	fmt.Println(r.MatchString("peach"))

	// А тут, ми знаходимо збіги для нашого regexp'а.
	fmt.Println(r.FindString("peach punch"))

	// Цей приклад теж знайде збіги, але поверне початковий
	// та кінцевий індекси, замість рядка що збігається.
	fmt.Println(r.FindStringIndex("peach punch"))

	// Варіації методів `Submatch` включать інформацію про
	// повні збіги та часткові збіги в межах повних. Наприклад -
	// це поверне інформацію про regexp'и `p([a-z]+)ch`
	// та `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Схожим чином - ми можемо отримати індекс повного і
	// часткового збігу.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Варіації методів `All` можуть застосовуватись до усіх збігів,
	// а не лише - до першого. Приклад знаходження усіх збігів.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Варіації `All` також доступний і іншим методам (наприклад
	// тим, що ми вже бачили).
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Другий аргумент (заданий додатнім цілим числом), допомагає
	// обмежити кількість збігів.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Приклади що ми розглянули, отримують аргумент типу рядок і
	// використовують імя на кшталт `MatchString`, але ми також можемо
	// передати зріз байтів - `[]byte` (та прибрати `String`
	// з імені методу) і отримаємо аналогічний метод, що буде
	// працювати із зрізами байтів.
	fmt.Println(r.Match([]byte("peach")))

	// Створюючи константи regexp'ів - ви можете
	// використати метод `MustCompile`, замість `Compile`.
	// Простий `Compile` не працює з константами - оскільки повертає
	// два значення, a `MustCompile` [панікує](./panic) у разі проблеми
	// що робить його безпечним для використання під час компіляції.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Пакет `regexp` можна використовувати для заміни часткових
	// збігів рядків іншими значеннями.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Варіація  методів `Func` дозволяє трансформувати текст за допомогою функції аргументу.
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
idx: [0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
regexp: p([a-z]+)ch
a <fruit>
a PEACH

# Радимо, зачінити документацію пакету
# [`regexp`](http://golang.org/pkg/regexp/).
```
