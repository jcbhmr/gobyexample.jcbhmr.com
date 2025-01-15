# Цикли|For
```go
// `for` це єдина конструкція циклу в Go. Нижче наведено
// три базові типи для циклів `for`.

package main

import "fmt"

func main() {

	// Найбільш розповсюдженим є тип з єдиною інструкцією.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Класичний ініціалізація/умова/післяумова `for` цикл,
	// також відомий як С-style `for`.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` без умови буде відбуватись постійно аж
	// допоки програма не перерве його, за допомогою
	// ключових слів `break` (що перериває цикл) або
	// `return` (що повертає значення з функції).
	for {
		fmt.Println("loop")
		break
	}

	// Існує також ключове слово `continue` - призначенням
	// якого є переходу до наступної ітерації циклу.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```
```sh
$ go run for.go
1
2
3
7
8
9
loop
1
3
5

# Ми ще побачимо інші форми `for` пізніше
# коли ми будемо працювати з твердженями `range`,
# каналами та іншими структурами даних.
```