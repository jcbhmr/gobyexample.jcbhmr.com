# Ошибки (Errors)|Errors
```go
// В Go принято сообщать об ошибках через явное, отдельное
// возвращаемое значение. Это контрастирует с исключениями,
// используемыми в таких языках, как Java и Ruby, и
// перегруженным одиночным значением результата/ошибки,
// иногда используемым в подходе C. Go, позволяет легко
// увидеть, какие функции возвращают ошибки, и обрабатывать
// их, используя те же языковые конструкции, которые
// используются для любых других задач без ошибок.

package main

import (
	"errors"
	"fmt"
)

// По соглашению, ошибки - это последнее возвращаемое
// значение с типом `error` в стандартной реализации.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` создает стандартную `ошибку` с
		// указаннным сообщением
		return -1, errors.New("can't work with 42")

	}

	// Значение `nil` в поле ошибки, говорит о том, что
	// ошибок нет.
	return arg + 3, nil
}

// Можно использовать пользовательские типы в качестве
// ошибок, применяя к ним метод `Error()`. Вот вариант
// в примере выше, который использует пользовательский
// тип для явного представления ошибки аргумента.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// В этом случае мы используем синтаксис `&argError`
		// для создания новой структуры, предоставляя
		// значения для двух полей `arg` и `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// Два цикла ниже тестируют каждую из наших функций,
	// возвращающих ошибки. Обратите внимание, что
	// использование встроенной проверки ошибок в строке
	// `if` является обычный явлением в Go.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Если вы хотите использовать данные в пользовательской
	// ошибке, вам нужно получить ошибку как экземпляр
	// пользовательского типа через утверждение типа.
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
```
```sh
$ go run errors.go
f1 worked: 10
f1 failed: can't work with 42
f2 worked: 10
f2 failed: 42 - can't work with it
42
can't work with it

# Посмотрите [эту статью](http://blog.golang.org/2011/07/error-handling-and-go.html)
# в блоге Go для более подробной информации об ошибках.
```
