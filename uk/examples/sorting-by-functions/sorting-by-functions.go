// Інколи виникає необхідність у реалізації сортування,
// але по-іншому, не в натуральному порядку. Наприклад,
// сортування рядків за довжиною, на противагу сортуванню
// за алфавітом. Спробуємо розібратись як це реалізувати у GO.

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// Ми імплементуємо функцію порівняння довжини рядків.
	// `cmp.Compare` нам дуже допоможе.
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// Тепер ми викличемо `slices.SortFunc` з нашою
	// функцією порівняння.
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// Ви можете використовувати туж техніку для сортування зрізу
	// значень що не є базовими типами.
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "Jax", age: 37},
		{name: "TJ", age: 25},
		{name: "Alex", age: 72},
	}

	// Сортування `people` за віком використовувючи `slices.SortFunc`.
	//
	// Зауваження: якщо `Person` структура велика,
	// спробуйте скористатись зрізом що містить `*Person`
	// і змінити сортуючу функцію відповідно. Якщо в сумнівах [вимірюйте щвидкодію](testing-and-benchmarking)!
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
