# Logging

```go,editable
// The Go standard library provides straightforward
// tools for outputting logs from Go programs, with
// the [log](https://pkg.go.dev/log) package for
// free-form output and the
// [log/slog](https://pkg.go.dev/log/slog) package for
// structured output.
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// Simply invoking functions like `Println` from the
	// `log` package uses the _standard_ logger, which
	// is already pre-configured for reasonable logging
	// output to `os.Stderr`. Additional methods like
	// `Fatal*` or `Panic*` will exit the program after
	// logging.
	log.Println("standard logger")

	// Loggers can be configured with _flags_ to set
	// their output format. By default, the standard
	// logger has the `log.Ldate` and `log.Ltime` flags
	// set, and these are collected in `log.LstdFlags`.
	// We can change its flags to emit time with
	// microsecond accuracy, for example.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// It also supports emitting the file name and
	// line from which the `log` function is called.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// It may be useful to create a custom logger and
	// pass it around. When creating a new logger, we
	// can set a _prefix_ to distinguish its output
	// from other loggers.
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// We can set the prefix
	// on existing loggers (including the standard one)
	// with the `SetPrefix` method.
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// Loggers can have custom output targets;
	// any `io.Writer` works.
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// This call writes the log output into `buf`.
	buflog.Println("hello")

	// This will actually show it on standard output.
	fmt.Print("from buflog:", buf.String())

	// The `slog` package provides
	// _structured_ log output. For example, logging
	// in JSON format is straightforward.
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// In addition to the message, `slog` output can
	// contain an arbitrary number of key=value
	// pairs.
	myslog.Info("hello again", "key", "val", "age", 25)
}
```

```sh
# Sample output; the date and time
# emitted will depend on when the example ran.
$ go run logging.go
2023/08/22 10:45:16 standard logger
2023/08/22 10:45:16.904141 with micro
2023/08/22 10:45:16 logging.go:40: with file/line
my:2023/08/22 10:45:16 from mylog
ohmy:2023/08/22 10:45:16 from mylog
from buflog:buf:2023/08/22 10:45:16 hello

# These are wrapped for clarity of presentation
# on the website; in reality they are emitted
# on a single line.
{"time":"2023-08-22T10:45:16.904166391-07:00",
 "level":"INFO","msg":"hi there"}
{"time":"2023-08-22T10:45:16.904178985-07:00",
	"level":"INFO","msg":"hello again",
	"key":"val","age":25}
```
