// [_Хеши SHA1_](http://en.wikipedia.org/wiki/SHA-1) часто
// используются для вычисления коротких идентификаторов
// для двоичных или текстовых BLOB-объектов. Например,
// [система контроля версий git](http://git-scm.com/) широко использует SHA1 для
// идентификации версионных файлов и каталогов. Вот как
// вычислить хэши SHA1 в Go.

package main

// Go реализует несколько хеш-функций в пакете
// `crypto/*`.
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// Для генерации хеша используется функция `sha1.New()`,
	// `sha1.Write(bytes)`, затем `sha1.Sum([]byte{})`.
	// Тут мы создаем новый хеш.
	h := sha1.New()

	// В `Write` необходимо передать байты. Если у вас есть
	// строка `s`, используйте `[]byte(s)`, чтобы привести
	// ее к байтам.
	h.Write([]byte(s))

	// Получаем окончательный результат хеш-функции в виде
	// байтового фрагмента. Аргумент в `Sum` добавляет к
	// существующему фрагменту байты, но обычно он не используется.
	bs := h.Sum(nil)

	// Значения SHA1 часто печатаются в шестнадцатеричном формате,
	// например, в git commit. Используйте `%x` для преобразования
	// результатов хеширования в шестнадцатеричную строку.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
