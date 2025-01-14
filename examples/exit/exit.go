// Используйте `os.Exit` для немедленного выхода с
// полученныем статусом.

package main

import (
	"fmt"
	"os"
)

func main() {

	// `defer` _не будет_ запускаться при использовании
	// `os.Exit`, поэтому этот `fmt.Println` никогда не
	// будет вызываться.
	defer fmt.Println("!")

	// Выход со статусом 3.
	os.Exit(3)
}

// Обратите внимание, что в отличие, например, от C,
// Go не использует целочисленное возвращаемое значение
// из `main`, чтобы указать состояние выхода. Если
// вы хотите выйти с ненулевым статусом, вы должны
// использовать `os.Exit`.
