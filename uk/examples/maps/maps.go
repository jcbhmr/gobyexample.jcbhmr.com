// _Мапи_ це [асоційований тип даних](http://en.wikipedia.org/wiki/Associative_array) в Go
// (в інших мовах програмування, такі структури даних ще називають: _хеш таблиця_ або _словник_).

package main

import (
	"fmt"
	"maps"
)

func main() {

	// Скориставшись функцією [`make`](./slices):
	// `make(map[key-type]val-type)`, ми ініціюємо створення мапи.
	m := make(map[string]int)

	// Задати/Створити значенням можна за допомогою,
	// вже знайомого, синтаксису доступу за індексом,
	// де індекс це `ключ`,
	// наприклад - <nobr>`name[key] = val`<nobr>.
	m["k1"] = 7
	m["k2"] = 13

	// Друкуючи мапу з `fmt.Println` ми побачимо усі існуючи
	// пари ключ/значення, які належить мапі.
	fmt.Println("map:", m)

	// Отримати значення за ключем можна звернувшись до
	// мапи за індексом `name[key]`.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Якщо ключа не існує,
	// [нульове значення](https://go.dev/ref/spec#The_zero_value)
	// типу буде повернуто.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// Функція `len` повертає кількість пар у мапі.
	fmt.Println("len:", len(m))

	// Функція `delete` видаляє пару ключ/значення з мапи.
	delete(m, "k2")
	fmt.Println("map:", m)

	// Щоб видалити *усі* пари ключ/значення з карти, скористайтесь
	// `clear`.
	clear(m)
	fmt.Println("map:", m)

	// Необов'язкове друге значення, отримане під час запиту значення
	// за ключем з мапи, покаже чи присутній даний ключ в мапі.
	// Це стає в нагоді для перевірки існування ключів ( особливо
	// ключів _з нульовим значенням_ (наприклад `0` або `""`)).
	// В цьому прикладі перевірки існування, нам непотрібно значення,
	// а отже, ми ігноруємо його за допомогою _пустого ідентифікатора_ `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Використовуючи наступний синтаксис ми можемо ініціалізувати мапу
	// одним рядком.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// Пакет `maps` містить корисні функції для мап.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}