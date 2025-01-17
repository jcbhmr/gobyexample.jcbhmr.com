# Функции с переменным числом аргументов (Variadic Functions)|Variadic Functions
```go
// [_Функции с переменным числом аргументов_](http://en.wikipedia.org/wiki/Variadic_function)
// могут быть вызваны с любым количество аргументов.
// Пример такой функции - это `fmt.Println`.

package main

import "fmt"

// Это функция, которая может принимать любое количество
// аргументов целых чисел.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Такие функции можно вызывать обычным способом.
	sum(1, 2)
	sum(1, 2, 3)

	// Если у вас есть несколько аргументов в срезе,
	// вы можете применить его к функции с переменным
	// числом аргументов таким образом `func(slice...)`.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```
```sh
$ go run variadic-functions.go 
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10

# Другим ключевым аспектом функций в Go являются
# замыкания, которые мы рассмотрим далее.
```
