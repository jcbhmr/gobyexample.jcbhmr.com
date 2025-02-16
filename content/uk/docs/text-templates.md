# Текстові Шаблони|Text Templates
```go
// У Go присутня підтримка для створення динамічного контенту або
// демонстрації користувацького контенту за допомогою пакунка `text/template`
// Пакет родич `html/template` дає теж API але має додаткові фічі щодо безпеки
// які повинні бути використані для генерації HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// Ми можемо стоврити новий шаблон і відпарсити його тіло з рядку.
	// Шаблони це мікс статичного тексту та "екшенів" закритих у
	// `{{...}}` що використовуються для динамічного втавляння контенту.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Альтернативно, ми можемо скористатись `template.Must` функцією
	// що панікує у разі повернення помилки `Parse`. Це супер корисно
	// для шаблонів ініціалізованих у глобальному просторі..
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Під час виконанням шаблону ми генеруємо його тест з
	// специфічними даними його дій. `{{.}}` буде замінено
	// значенням переданим в `Execute`.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Функція помічник якою ми скористаємось далі.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Якщо данні це структура, ми можемо скористатись `{{.FieldName}}`
	// для доступу до поля. Поля мають бути експортними (публічними)
	// на момент виконання шаблону.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Це є відноситься до мап; з мапами немає обмежень щодо регістру ключів.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else дають нам можливість умовного виконання шаблону. Значення вважається
	// false якщо це нульове значення типу, наприклад 0, пуста строка, nil pointer,
	// абощо.
	// Цей приклад показує іншу фічу шаблонів екшен `-` (мінус) в дії прибирання пробілів.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// длоки діапазону дають нам можливість циклу по зрізам, мапам, каналам чи масивам,
	// Всередині `{{.}}` встановлено поточний елемент ітерації.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
```
```sh
$ go run templates.go 
Value: some text
Value: 5
Value: [Go Rust C++ C#]
Name: Jane Doe
Name: Mickey Mouse
yes 
no 
Range: Go Rust C++ C# 
```
