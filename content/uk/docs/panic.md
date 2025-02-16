# Паніка|Panic
```go
// `panic` або `паніка` зазвичай означає те, що щось пішло
// неочікувано погано. Дуже погано. В більшості випадків, ми
// використовуємо її задля швидкого провалу
// на помилках які зазвичай не мали б виникати, або
// на помилках які ми не здатні опрацювати акуратно.

package main

import "os"

func main() {

	// Ми скористаємось панікою в цьому прикладі для індикації
	// неочікуваної помилки. Єдина мета цієї ділянки коду
	// запанікувати.
	panic("маємо проблему")

	// Загальне використання `panic`и - перервати виконання функції,
	// воно повертає помилку яку ми не знаємо як опрацьовувати
	// (або не хочемо). Ось приклад `panic`и де ми отримали
	// неочікувану помилку, під час створення файлу, з якою ми не
	// дуже хочемо мати справу.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
```sh
# Запуск цієї програми призведе до її паніки, виводу
# повідомлення помилки та сліду горутини, та виходу з
# не нульовим статусом.

# When first panic in `main` fires, the program exits
# without reaching the rest of the code. If you'd like
# to see the program try to create a temp file, comment
# the first panic out.
$ go run panic.go
panic: маємо проблему

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Зауважте що, на відміну від інших мов що
# використовують виключення (exceptions) для опрацювання
# помилок, в GO ідіоматично викорисовувати -
# _повернення помилок_, що вказуватимуть
# на проблеми як умога частіше.
```
