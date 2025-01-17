# Массивы (Arrays)|Arrays
```go
// В Go, _массив_ это числовой ряд элементов определенной
// длины.

package main

import "fmt"

func main() {

	// В данном примере мы создаем массив `a`, который
	// содержит 5 элементов с типом `int`. Тип элементов
	// и длина являются частью типа массива. По-умолчанию
	// массив заполняется нулевыми значениями, например,
	// в случае `int` нулевое значение - 0.
	var a [5]int
	fmt.Println("emp:", a)

	// Мы можем установить значение по индексу элемента
	// следующим образом:`array[index] = value`.
	// Получить значение можно аналогично - `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Встроенная функция `len` возвращает длину массива.
	fmt.Println("len:", len(a))

	// Так можно инициалзировать и заполнить массив
	// значениеми в одну строку
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Тип `массив` является одномерным. Но вы можете
	// совмещать типы, для создания многомерных
	// структур.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
```
```sh
# Обратите внимание, что массивы отображаются в виде
# [v1 v2 v3 ...] при выводе с помощью fmt.Println.
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d:  [[0 1 2] [1 2 3]]

# В Go вы будете встречать _срезы_ гораздо чаще, чем
# массивы. Срезы рассмотрим далее
```
