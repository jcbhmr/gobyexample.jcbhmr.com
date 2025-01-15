# Тестування та Вимірювання|Testing and Benchmarking
```go
// Unit-тестування (або модульне тестування) є важливою частиною написання
// програм мовою Go. Пакет `testing` надає інструменти,
// необхідні для написання таких тестів, а команда `go test`
// дозволяє їх виконання.

// Для демонстрації - приведемо код пакету
// `main` (але підійде будь-який пакет). Код тестів
// зазвичай розміщують у тому ж самому пакеті,
// код якого тестують.
package main

import (
	"fmt"
	"testing"
)

// Ми протестуємо просту реалізацію знаходження
// найменшого числа. Зазвичай, якщо код, який ми збираємось
// тестувати знаходиться у файлі з іменем, наприклад `intutils.go`,
// то файл з кодом тестів має називатись відповідно `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Тест створюється шляхом написання функції з іменем,
// що починається з `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` повідомить про помилку, але продовжить
		// виконання тесту. `t.Fail*` повідомить про помилку
		// та негайно зупинить виконання тесту.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Написання тестів може бути повторюваним, тому ідіоматично в Go,
// використовувати *табличні тести* (*table-driven tests*),
// де тестові - вхідні та очікувані дані наведені в таблиці, та
// в одному циклі проходити ці дані й виконувати тести.
func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		// t.Run дозволяє запускати "підтести", по одному для кожного
		// запису таблиці. Вони відображаються окремо
		// при виконанні `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Вимірювальні тести зазвичай знаходяться у `_test.go` файла
// і називаються з префіксом `Benchmark`. Тест виконує кожне
// вимірювання function кілька разів, збільшуючи `b.N` кожної
// ітерації  допоки не назбирає точні вимірювання.
func BenchmarkIntMin(b *testing.B) {
	// Типово вимірювання запускають функцію що ми вимірюємо у
	// циклі `b.N` разів.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
```
```sh
# Запуск усіх тестів у поточному пакеті у
# режимі "докладно".
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok  	examples/testing-and-benchmarking	0.023s

# Запуск усіх вимірювань у поточному проекті. Усі тести
# запустились до вимірювань, Прапорець `bench`
# відфільтровує імена функцій за допомогою регулярного
# виразу
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: examples/testing
BenchmarkIntMin-8 1000000000 0.3136 ns/op
PASS
ok  	examples/testing-and-benchmarking	0.351s
```
