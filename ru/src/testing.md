# Тестирование (Testing)|Testing
```go
// Unit тестирование является важной частью написания
// правильных программ Go. Пакет тестирования `testing`
// предоставляет инструменты, необходимые для
// написания unit тестов, а команда `go test` запускает
// тесты.

// Для наглядности этот код находится в `main` пакете,
// но это может быть любой пакет. Тестовый код обычно
// находится в том же пакете, что и код, который он тестирует.
package main

import (
	"fmt"
	"testing"
)

// Мы будем тестировать эту простую реализацию
// целочисленного минимума. Обычно код, который мы
// тестируем, находится в исходном файле с именем
// что-то вроде `intutils.go`, а тестовый файл для
// него будет называться `intutils_test.go`.
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Тест создается путем написания функции с именем,
// начинающимся с `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` сообщит об ошибках теста, но продолжит
		// выполнение теста. `t.Fail*` сообщит об ошибках
		// теста и немедленно остановит тест.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Написание тестов может быть повторяющимся, поэтому
// идиоматично использовать *table-driven style*, где тестовые
// входы и ожидаемые выходы перечислены в таблице, а один
// цикл проходит по ним и выполняет тестовую логику.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
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
		// t.Run позволяет запускать «подтесты», по одному
		// для каждой записи таблицы. Они показываются
		// отдельно при выполнении `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
```
```sh
# Run all tests in the current project in verbose mode.
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
ok  	examples/testing	0.023s
```
