# Варіативні Функції|Variadic Functions
```go
// [_Варіативна функція_](http://en.wikipedia.org/wiki/Variadic_function)
// може бути викликана з будь-якою кількістю елементів.
// Так-от, `fmt.Println` - є прикладом варіативної функції.

package main

import "fmt"

// Ця функція, прийме своїм аргументом довільну кількість цілих чисел.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	// Within the function, the type of `nums` is
	// equivalent to `[]int`. We can call `len(nums)`,
	// iterate over it with `range`, etc.
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Варіативні функції, можуть бути викликані звичайним
	// чином з окремо вказаними аргументами.
	sum(1, 2)
	sum(1, 2, 3)

	// Також можливо 'розпакувати' аргументи, які знаходяться вигляді [`зрізу`](./slices),
	// скориставшись синтаксисом розпаковування - `func(slice...)`.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```
```sh
$ go run variadic-functions.go
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

# Іншим ключовим аспектом функцій в Go - є їх
# здатність створювати замикання, що є
# нашою наступною темою.
```
