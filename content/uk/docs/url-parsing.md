# Розбір URL|URL Parsing
```go
// _Уніфікований локатор ресурсів_, більш відомий
// за своєю англомовною абревіатурою - URL, надає змогу
// визначати [локацію ресурсі](https://adam.herokuapp.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/) у інтернеті (і не тільки). Використовуючи Go - ми маємо змогу розбити URL на
// частинки та отримати їх автоматично.
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// Наш приклад представляє собою URL, що містить: схему,
	// інформацію аутентифікації, хост, порт, шлях,
	// параметри запиту і фрагмент запиту.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Метод розбору `url.Parse` також повертає "[помилку](./errors)", що
	// дозволяє перевірити - чи виникли проблеми під час розбору URL.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Отримати доступ до "Схеми" (і інших складових) дуже просто.
	fmt.Println(u.Scheme)

	// `User` зберігає усю інформацію щодо аутентифікації;
	// Зверніться до "під-полів" `Username` та `Password` для отримання
	// індивідуальних значень.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// `Host` складається як з імені хосту, так і порту,
	// якщо він присутній. Скористайтесь `SplitHostPort`
	// для отримання складових значення.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// А таким чином ми звернемось до `шляху` та `фрагмента` після `#`.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// Щоб отримати параметри запиту в рядку у форматі
	// `k=v` скористайтесь полем структури URL - `RawQuery`.
	// Також можливо - розкласти парметри запиту у мапу. Р
	// Самі параметри запиту збережені у мапі у вигляді `ключ-рядок`
	// де значення зберігається в зрізі рядків, отож звертайтесь до значення
	// за індексом `[0]` (якщо вам потрібно буде тільки перше значення).
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
```
```sh
# Запустіть нашу програму, і отримайте
# розібраний на шматочки URL.
$ go run url-parsing.go
postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
```
