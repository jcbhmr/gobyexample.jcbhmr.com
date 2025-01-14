// Інколи, необхідно запускати інші процеси.

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// Розпочнемо ми з простої команди, яка не приймає
	// жодних аргументів або вводу, і яка лише друкує "щось"
	// до `stdout`. Функція `exec.Command` створить об'єкт
	// який представляє собою цей зовнішній процес.
	dateCmd := exec.Command("date")

	// `Output` - інший метод, який бере на себе управління
	// загальними випадками запуску процесу, очікує його
	// завершення та отримання виводу. Якщо помилок не було,
	// [зріз](slices) `dateOut` буде наповнено байтами відповіді
	// процесу `date`, який ми викликали.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// `Output` та інші методи `Command` повернуть
	// `*exec.Error` якщо спостерігалась помилка на
	// стороні додатку що запускається в окремому процесі.
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// Тепер ми готові розглянути більш просунутий приклад,
	// де ми перенаправимо дані у [`потік введення`](https://uk.wikipedia.org/wiki/Стандартні_потоки#Стандартне_введення) зовнішнього процесу та
	// отримаємо результати з його [`потоку виведення`](https://uk.wikipedia.org/wiki/Стандартні_потоки#Стандартне_введення).
	grepCmd := exec.Command("grep", "hello")

	// Ми чітко вказуємо - "захопи перенаправлення вводу та виводу",
	// запускаємо процес, виконуємо необхідне введення інформації
	// до нього, читаємо результатний вивід і, нарешті, очікуємо
	// завершення процесу.
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Ми пропустили перевірку помилок у попередньому прикладі,
	// але ви можете використовувати звичайний прийом `if err != nil`
	// для цього. Ми вже скористались результатом отриманим з `StdoutPipe`,
	// у схожим чином, ви можете отримати відповідь що надходить з `StderrPipe`.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// Зауважимо - коли ми породжуємо процеси нам необхідно надати
	// чітко розділені команду і масив аргументів, напротивагу передачі
	// всієї команди записаної в один рядок. Якщо ж вам, таки, кортить
	// передати все одним рядком, скористайтесь `bash`'овим параметром `-c`:
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
