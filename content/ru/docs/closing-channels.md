# Закрытие каналов (Closing Channels)|Closing Channels
```go
// _Закрытие_ канала означает, что по нему больше не
// будет отправлено никаких значений. Это может быть
// полезно для сообщения получателям о завершении.

package main

import "fmt"

// В этом примере мы будем использовать канал `jobs` для
// передачи задания, которое должна быть выполнена из
// `main()` в горутине. Когда у нас больше не будет
// заданий для воркера, мы `закроем` канал `jobs`.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Вот наш воркер. Он многократно получает из канала
	// `jobs` значения `j, more := <-jobs`. В этой специальной
	// форме получения с двумя значениями `more` значение
	// будет `ложным`, если `jobs` были `закрыты`, а все
	// значения в канале уже получены. Мы используем
	// это, чтобы уведомить о `выполнении`, когда мы
	// проработали все наши работы.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Отправляем 3 сообщения в канал `jobs`, и закрываем
	// его.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Мы ожидаем выполнения всех каналов используя
	// [синхронизацию](channel-synchronization), рассмотренную
	// нами ранее.
	<-done
}
```
```sh
$ go run closing-channels.go 
sent job 1
received job 1
sent job 2
received job 2
sent job 3
received job 3
sent all jobs
received all jobs

# Идея закрытых каналов естественно приводит нас к
# следующему примеру: `range` по каналам.
```
