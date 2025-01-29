# Пулы воркеров (Worker Pools)|Worker Pools
```go
// В этом примере мы рассмотрим, как реализовать
// _пул воркеров_ с использованием каналов и
// горутин.

package main

import (
	"fmt"
	"time"
)

// Это воркер, который мы будем запускать в нескольких
// параллельных инстансах. Эти воркеры будут получать
// задания через канал `jobs` и отсылать результаты
// в `results`. Мы будем ожидать одну секунду для
// каждого задания для имитации тяжелого запроса.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Чтобы использовать наш воркер пул, нам нужно
	// отправить им задание и получить результаты выполнения.
	// Для этого мы делаем 2 канала.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Стартуем 3 воркера, первоначально заблокированных,
	// т.к. еще нет заданий.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Посылаем 5 `заданий (jobs)` и затем `закрываем`
	// канал, сообщая о том что все задания отправлены.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Наконец мы собираем все результаты. Это также
	// гарантирует, что горутины закончились. Альтернативный
	// способ ожидания нескольких процедур заключается
	// в использовании [WaitGroup](waitgroups).
	for a := 1; a <= 5; a++ {
		<-results
	}
}
```
```sh
# Наша программа показывает 5 заданий, выполняемых
# различными воркерами. Программа занимает всего около
# 2 секунд, несмотря на выполнение около 5 секунд общей
# работы, потому что одновременно работают 3 воркера.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
```
