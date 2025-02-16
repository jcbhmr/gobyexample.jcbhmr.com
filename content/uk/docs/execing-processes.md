# Заміна поточного процесу новим|Exec'ing Processes
```go
// Тепер, коли ми розглянули
// [запуск сторонніх процесів](spawning-processes), ми знаємо, що ми це робимо
// коли необхідно, щоб запущений процес був доступний нашому Go процесу.
// Інколи, все чого ми бажаємо - це лише повністю замінити поточний Go-процес
// іншим (можливо, навіть - не "Go-процесом"), для цього ми скорастаємось реалізацією
// одного з системних викликів - [exec](http://en.wikipedia.org/wiki/Exec_(operating_system)).

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// Спробуємо виконати `ls` у нашому прикладі. Go необхідно абсолютний
	// шлях до програми, яку ми хочемо виконати - отож ми скористаємось
	// `exec.LookPath` щоб знайти її (і можливо ми знайдемо її як `/bin/ls`).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Параметри передаються в `Exec` у вигляді [`зрізу`](slices)
	// (напротивагу одному рядку). Ми забезпечимо `ls`
	// кількома звичайними (для цієї команди) аргументами.
	args := []string{"-h", "-o"}

	// `Exec` також потребує [змінні оточення](environment-variables)
	// і ми ж надамо йому наше поточне середовище.
	env := os.Environ()

	// А ось - і власне виклик `syscall.Exec`. Якщо виклик успішний,
	// виконання нашого процесу завершиться і буде замінено процесом `/bin/ls -h -o`,
	// а якщо виникне помилка - то ми отримаємо [паніку](panic) з текстом помилки.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
```
```sh
# Виконання нашої програми замінено виконанням `ls`.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Note that Go does not offer a classic Unix `fork`
# function. Usually this isn't an issue though, since
# starting goroutines, spawning processes, and exec'ing
# processes covers most use cases for `fork`.
```
