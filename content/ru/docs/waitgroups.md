# WaitGroups|WaitGroups
```go
// Для ожидания выполнения нескольких горутин, мы можем
// использовать встроенную конструкцию *WaitGroup*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Эта функция, которую мы будем запускать в каждой
// горутине. Обратите внимание, что WaitGroup должна
// быть передана в функцию по указателю.
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep симулирует тяжелую задачу.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	// Оповестить WaitGroup что воркер выполнился
	wg.Done()
}

func main() {

	// Эта WaitGroup используется для ожидания выполнения
	// всех горутинё запущенных здесь.
	var wg sync.WaitGroup

	// Запускаем несколько горутин и инкрементируем счетчик
	// в WaitGroup для каждой запущенной горутины.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Блокируем звершение программы до момента, пока
	// счетчик WaitGroup снова не станет равным 0.
	// Это будет означать, что все горутины выполнились.
	wg.Wait()
}
```
```sh
$ go run waitgroups.go
Worker 5 starting
Worker 3 starting
Worker 4 starting
Worker 1 starting
Worker 2 starting
Worker 4 done
Worker 1 done
Worker 2 done
Worker 5 done
Worker 3 done

# Порядок воркеров начинающихся и выполненных, вероятно
# будет изменяться при каждом запуске.
```
