# Групи Очікування (WaitGroups)|WaitGroups
```go
// Щоб очікувати на закінчення роботи кількох горутин,
// ми можемо скористатись *групою очікування*.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Це функція, яку ми виконаємо в кожній горутині.
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// Sleep to simulate an expensive task.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// Ми скористаємось `WaitGroup` для очікування результату роботи
	// усіх горутин.
	// Зауважте, що WaitGroup має бути передано до функції - як вказівник.
	var wg sync.WaitGroup

	// Збільшимо лічильник Групи Очікування для кожної запущеної
	// горутинки.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Уникайте перевикористання  значення `i` для кожного замикання.
		// Дивіться [FAQ](https://golang.org/doc/faq#closures_and_goroutines)
		i := i

		// Заверніть виклик воркеру у замикання що переконається доповісти
		// WaitGroup про завершення роботи. при цьомі підходи воркер не має потреби
		// бути вкурсі про якісь примітиви багатопоточного виконання.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Група Очікування блокує виконання до моменту поки лічильник
	// знову не буде відкатане до позначки 0.
	wg.Wait()

	// Зауважте що данний підхід не має визначеного шляху
	// передавати помилки з воркерів. Для більш докладного використання,
	// зверніть увагу на
	// [пакунок `errgroup`](https://pkg.go.dev/golang.org/x/sync/errgroup).
}
```
```sh
# Порядок старту та завершення буде іншим.
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
```
