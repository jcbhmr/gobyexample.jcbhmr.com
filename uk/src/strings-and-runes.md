# Рядки та руни|Strings and Runes
```go
// Рядок у GO це зріз байтів лише для читання (не-мутабельний).
// Мова та стандартна бібліотека  відносяться до
// рядків спеціально як до контейнерів тексту закодованого у
// [UTF-8](https://en.wikipedia.org/wiki/UTF-8).
// У інших мовах, рядки збудовані з "символів" - у Go ж концеп символу
// називається руною це ціле цисло що представляє собою номер utf8
// коду.
// [Цей запис у блозі Go](https://go.dev/blog/strings) непогане введедння
// в тему.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// `s` це рядок до якого ми додаємо літеральне значення
	// що означає слово "Привіт" у Тайській мові. Строкові літерали
	// закодовано UTF-8.
	const s = "สวัสดี"

	// Оскільки рядки це еквівалень `[]byte`, довжина тепер буде
	// довжиною байтів що містить рядок.
	fmt.Println("Len:", len(s))

	// ІНдексування рядку дає нам байти що знаходяться за кожним індексом.
	// Цей цикл генерує щістнадцятирічні значення усів байтів що конструюють
	// utf8 символи (код поніти) e `s`.
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// Щоб порахувати скільки _rune_ у ряду, ми скористаємось
	// пакетом `utf8`. Зауважимо що час виконання
	// `RuneCountInString` залежить від розміру строки,
	// тому що вона має декодувати кожену UTF-8 руну послідовно.
	// А деякі такйські символи являють собою численні UTF-8
	// символи, отож результат може бути цікавим.
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Цикл `range` підтримує рядки окремо і розкодовує кожну
	// `rune` разом з її індексами (офсетами) у рядку.
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Ми можемо досягти того ж резльтату використовуючи
	// функцію `utf8.DecodeRuneInString` напряму.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Це демонстрація передачі значення руни у функцію.
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Значення замкнені у одинарні лапки є _літералами руни_.
	// Ми можемо порівнювати значення руни до її літерала напряму.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
```
```sh
$ go run strings-and-runes.go
Len: 18
e0 b8 aa e0 b8 a7 e0 b8 b1 e0 b8 aa e0 b8 94 e0 b8 b5 
Rune count: 6
U+0E2A 'ส' starts at 0
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15

Using DecodeRuneInString
U+0E2A 'ส' starts at 0
found so sua
U+0E27 'ว' starts at 3
U+0E31 'ั' starts at 6
U+0E2A 'ส' starts at 9
found so sua
U+0E14 'ด' starts at 12
U+0E35 'ี' starts at 15
```
