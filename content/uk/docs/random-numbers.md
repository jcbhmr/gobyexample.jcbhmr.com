# Генерація випадкових чисел|Random Numbers
```go
// Пакет `math/rand` дозволяє реалізовувати генерацію
// [псевдовипадкових чисел](https://uk.wikipedia.org/wiki/Генератор_псевдовипадкових_чисел).

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Наприклад, `rand.Intn` повертає випадкове ціле число в
	// діапазоні від нуля до n (`0 <= n < 100`).
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// А `rand.Float64` поверне випадкове число з плаваючою комою
	// (тип `float64`) в діапазоні, від нуля до 1-ці (`0.0 <= f < 1.0`).
	fmt.Println(rand.Float64())

	// Це може стати корисним для генерації чисел в діапазонах,
	// наприклад `5.0 <= f' < 10.0`.
	rand1 := (rand.Float64() * 5) + 5
	rand2 := (rand.Float64() * 5) + 5
	fmt.Println(rand1, ",", rand2)

	// Станадартний генератор чисел детермінований, отже він буде
	// видавати тіж самі послідовності чисел кожного разу.
	// Щоб отримати послідовності, які будуть мінятись, нам необхідно
	// надати початкове значення, або "seed" (як його ще називають).
	// Зауваження - це не дуже безпечно використовувати випадкові
	// числа які мають триматись в секреті, скористайтесь
	// пакетом `crypto/rand` якщо ваша ціль генерація більш
	// захищених даних.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Виклик результатного `rand.Rand` проходить аналогічно
	// функціям пакету `rand`.
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// Якщо початковим значенням буде одне і теж число
	// результат, теж, буде однаковий для різних викликів генераторів.
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
```
```sh
# Depending on where you run this sample, some of the
# generated numbers may be different. Note that on
# the Go playground seeding with `time.Now()` still
# produces deterministic results due to the way the
# playground is implemented.
$ go run random-numbers.go
81,87
0.6645600532184904
7.123187485356329,8.434115364335547
0,28
5,87
5,87


# Щоб отримати довідку по наявним методом
# генерації випадкових чисел - зверніть увагу
# на документацію пакету [`math/rand`](http://golang.org/pkg/math/rand/).
```
