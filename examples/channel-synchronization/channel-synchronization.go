// Мы можем использовать каналы для синхронизации
// выполнения между горутинами. Вот пример
// использования блокирующего получения для ожидания
// завершения работы горутины. При ожидании завершения
// нескольких процедур вы можете использовать [WaitGroup](waitgroups).

package main

import (
	"fmt"
	"time"
)

// Эту функцию мы будем запускать в горутине. Канал
// `done` будет использован для оповещения другой
// горутины о том, что функция выполнена успешно.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Отправить значение, чтобы сказать что функция
	// выполнена успешно.
	done <- true
}

func main() {

	// Запустите воркера в горутине и передайте ему канал
	// для оповещения.
	done := make(chan bool, 1)
	go worker(done)

	// Блокируйте, пока мы не получим уведомление от
	// воркера из канала.
	<-done
}
