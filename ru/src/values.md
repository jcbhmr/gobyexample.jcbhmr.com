# Типы данных (Values)|Values
```go
// В Go существуют различные типы значений: строки,
// целые числа, числа с плавающей точкой, булевы и т.д.
// Вот некоторые примеры

package main

import "fmt"

func main() {

	// Строки могут быть сложены с помощью символа `+`.
	fmt.Println("go" + "lang")

	// Целый числа и числа с плавающей точкой.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Логические значения с логическими операторами
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```
```sh
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```