// Go's `slices` реалізує сортування для вбудованих
// та визначених користувачами типів. Спочатку познайомимось з
// сортуванням вбудованих типів.

package main

import (
	"fmt"
	"slices"
)

func main() {
	// Функції сортування абстрактні, і працюватимуть для будь-якого
	// _впорядкованого_ типу. За списком Сортваних типів завітайте
	// на сторінку [cmp.Ordered](https://pkg.go.dev/cmp#Ordered).
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Сортування цілих чисел - не дуже відрізняється від
	// сортування рядків.
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// Ми, також, можемо використати `sort` щоб перевірити
	// чи зріз вже відсортований.
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
