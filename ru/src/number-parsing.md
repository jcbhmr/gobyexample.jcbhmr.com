# Парсинг чисел (Number Parsing)|Number Parsing
```go
// Парсинг чисел из строк является распространенной задачей;
// Вот как это реализовать в Go.

package main

// Для решения этой задачи подойдет встроенный пакет
// `strconv`.
import (
	"fmt"
	"strconv"
)

func main() {

	// С помощью `ParseFloat`, параметр `64` говорит о том,
	// сколько битов точности необходимо использовать.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Для `ParseInt` 0 означает вывод базы из строки. `64`
	// необходимо, чтобы результат соответствовал 64 битам.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// `ParseInt` будет распознавать числа в шестнадцатеричной
	// системе.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` так же доступен.
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` это удобная функция для парсинга в десятеричный
	// `int`.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Функции парсинга возвращают ошибку в случае некорректных
	// аргументов.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
```
```sh
$ go run number-parsing.go 
1.234
123
456
789
135
strconv.ParseInt: parsing "wat": invalid syntax

# Далее мы рассмотрим парсинг URL.
```
