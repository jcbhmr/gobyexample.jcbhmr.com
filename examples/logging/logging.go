// Стандартна бібліотека Go надає як простиі засоби
// виводити логи програм за допомогою пакету [log](https://pkg.go.dev/log)
// для довільної форми логів так і пакет [log/slog](https://pkg.go.dev/log/slog)
// для стуктурованого логування.

package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {

	// Просто викликаючи функції типу `Println` пакету `log`
	// використає _standard_ логгер, який вже попередньо сконфігуровано
	// виводити до `os.Stderr`. Такі додаткові методи як  `Fatal*` або
	// `Panic*` припинять роботу програми після логування.
	log.Println("standard logger")

	// Логгер можливо сконфігуровати _прапорцями_ для встановленння
	// формату виводу. По-замовчуванню, стандартний логер має
	// `log.Ldate` та `log.Ltime` прапорці встановлено, вони збираються
	// `log.LstdFlags`. Ми можемо змінити ці прапорці, щоб, напряклад,
	// виводити час з мікросекундною акуратністю.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	//  Або виводити імя файлу з якого функція `log` виклиана.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// Може бути корисним створити власний логер, та
	// передавати його куди треба. Коли створбємо
	// власний логер, краще задавати префікс щоб відрізняти його
	// від інших.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// Ми можемо задати префікс і для існуючих логерів (включно з стандартним)
	// скориставшись методом `SetPrefix`.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Логерри можуть мати власні цілі логування; будь-який `io.Writer` підійде.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// Цей виклик запише лог до `buf`.
	buflog.Println("hello")

	// Виведемо наш лог вже у стандартний потік виводу.
	fmt.Print("from buflog:", buf.String())

	// Пакет `slog` надає структурований вивід логу.
	// Наприклад, форматування JSON для логів.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// На додачу до логів, `slog` може містити довільну кількість пар
	// ключ=значення.
	myslog.Info("hello again", "key", "val", "age", 25)
}
