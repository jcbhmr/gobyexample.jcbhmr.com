// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Вот пример форматирования времени в соответствии с
	// RFC3339 с использованием соответствующей константы.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Парсинг значений использует тот же синтаксис, что и `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Форматирование` и `парсинг` используют аргументы на
	// основе примеров. Обычно вы используете константу из пакета
	// `time`, но вы также можете предоставить собственные шаблоны.
	// Они должны использовать время в формате `Mon Jan 2
	// 15:04:05 MST 2006`, чтобы вывести на экран. Время в примере
	// должно быть точно таким, как показано: 2006 год, 15 для часа,
	// понедельник для дня недели и т.д.
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// Для числовых представлений вы также можете использовать
	// стандартное форматирование строки с методами времени.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` вернет ошибку. если не сможет распарсить данные.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}