// `//go:embed` is a [compiler
// directive](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives) that
// allows programs to include arbitrary files and folders in the Go binary at
// build time. Read more about the embed directive
// [here](https://pkg.go.dev/embed).
package main

// Імпортуйте пакунок `embed`; якщо ви не використовуєте жодного
// публічного ідентефікатора з пакунку, ви завжди можете зробити
// пустий імпорт за допомогою `_ "embed"`.
import (
	"embed"
)

// `embed` директива приймає шлях відносний директорії що містить
// Go файли. Ця директоива  вбудовує зміст файлу у
// змінну типу `string` що слідує за директивою.
//
//go:embed folder/single_file.txt
var fileString string

// Або вюудувати зміст файлу у зріз (`[]byte`).
//
//go:embed folder/single_file.txt
var fileByte []byte

// Ви такою можете імпортувати багато файлів або навіть папок
// за допомогою метасимволів. Таке використання потребує змінної типу
// [embed.FS type](https://pkg.go.dev/embed#FS), яка імплементує просту
// віртуальну файлову систему.
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {
	// Виводимо зміст `single_file.txt`.
	print(fileString)
	print(string(fileByte))

	// Виводимо зміст файлів з вбудованої папки.
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
