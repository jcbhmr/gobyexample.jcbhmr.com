# JSON
```go
// Go підтримує кодування та розкодування формату JSON
// на рівні, як базових типів даних, так і власних
// структур даних, так би кажучи "з коробки".

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Ми скористаємось цими структурами для демонстрації
// кодування та розкодування нетипових для Go структур даних.
type response1 struct {
	Page   int
	Fruits []string
}

// Тільки експортовані поля будуть закодовані/розкодовані у JSON.
// Поля повинні починатись з великої літери, щоб бути "експортованими".
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// Для початку ми розглянемо кодування базових
	// типів, доступних у Go, в формат JSON.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// В цьому прикладі ми працюємо з коллекціями
	// - зрізами та мапами, що будуть кодуватись в JSON
	// об'єкти передбачуваним (стандартизованим) чином.
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Пакет JSON автоматично кодує ваші типи даних.
	// Це дозволяє автоматично включати експортовані
	// поля до вашого виводу, або навпаки - виключати
	// не-експортовані з цього виводу. По-замовчуванню,
	// ці елементи будуть носити тіж назви, що і поля
	// структури яка кодувалась.
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Ви можете скористатись тегами або метаданим для
	// декларації полів у вашій структурі даних, гнучко налаштовуючи
	// логіку відображення або споживання даних за ключами.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Тепер, розглянемо як декодувати дані з формату JSON
	// у значення Go, на прикладі узагальненої структури.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Ми починаємо з того, що декларуємо змінну, якою пакет
	// JSON скористається для збереження розкодованих даних.
	// Змінна `dat` з типом `map[string]interface{}` буде зберігати
	// дані у вигляді мапи з рядковим ключем та значенням довільного типу.
	var dat map[string]interface{}

	// Розкодуємо наш JSON та проведемо перевірку на можливі помилки.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Для використання даних розкодованої мапи, нам необхідно буде
	// конвертувати (англійською `cast`) дані (тому що значення
	// нашої мапи - початково не типізовані - `interface{}`), щоб
	// продовжити роботу вже з чітко визначеним типом,
	// що підтримується Go. Для прикладу ми конвертуємо значення
	// мапи за ключем `num` в тип `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Доступ до вкладених даних - потребує серії конвертацій.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Слід використовувати ваші власні структури даних для
	// розкодування даних, це гарантія додаткової типо-безпеки,
	// що прибирає необхідність перевіряти типи під час використання.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// У, вже розглянутих, прикладах ми використовуємо зрізи байтів
	// та рядки як посередники між даними та JSON репрезентацією
	// для виводу до стандартного потоку даних. Подивіться, як
	// можна скористатись перенаправленням потоку за допомогою
	// власної реалізацію інтерфейсу `Writer`, наприклад `os.Writer`,
	// або навіть запис у HTTP Response.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
```
```sh
$ go run json.go
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}


# Ми ознайомились з базовою роботою Go з форматом
# JSON, і радимо прочитати статтю
# [JSON та Go](http://blog.golang.org/2011/01/json-and-go.html)
# та документацію [пакету JSON](http://golang.org/pkg/encoding/json/)
# для отримання більш розгорнутої інформації про сам
# пакет в Go.
```
