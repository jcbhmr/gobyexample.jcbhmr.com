 # Go за прикладом [![Build Status](https://travis-ci.org/butuzov/gobyexample.svg?branch=ukrainian)](https://travis-ci.org/butuzov/gobyexample)


Наповнення та інструментарій для [Go за Прикладом](https://butuzov.github.io/gobyexample/), сайту що навчає Go за допомогою анотованих прикладів.

### Загальне

Сайт "Go за прикладом" збудовано шляхом обробки коду та коментарів отриманих з першоджерельних файлів (що знаходяться в директорії `examples`) та форматуванню їх за допомогою шаблонів (з директорій `templates`) у статичні файли (що лежатимуть у директорії `public`). Інструменти що забезпечують весь процес створення сайт знаходяться у директорії `tools`, поряд з деякими залежностями (що лежать у `vendor`).


Створена директорія `public` може буде завантажена на будь-який shared/CDN/cloud хостинг.

### Побудова сайту

[![test](https://github.com/butuzov/gobyexample/actions/workflows/test.yml/badge.svg)](https://github.com/butuzov/gobyexample/actions/workflows/test.yml)

Щоб побудувати цей сайт вам потрібно буде Go. Запустіть:

```console
$ tools/build
```

Будувати в циклі:

```console
$ tools/build-loop
```

Обслуговувати сайт локально:

```
$ tools/serve
```

На відкрийте `http://127.0.0.1:8000/` у вашому браузері

### Зауваження щодо `public`

Ця директорія не оновлються в українській версії репозиторію. Перегенеруйте файли власноруч, у разі потреби.


### Publishing

Щоб завантажити сайт на AWS:

```bash
> gem install aws-sdk
> export AWS_ACCESS_KEY_ID=...
> export AWS_SECRET_ACCESS_KEY=...
> tools/upload
```

### Ліцензія

Ця робота є авторським правом Mark McGranaghan та ліцензована за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go's Гофер є авторським правом [Renée French](http://reneefrench.blogspot.com/) та ліцензовано за
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### Інші переклади

Переклади "Go by Example" від волонтерів доступні в наступних версіях:

* [Brazilian Portuguese](https://lcslitx.github.io/GoEmExemplos/) by [lcslitx](https://github.com/LCSLITX)
* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [Czech](http://gobyexamples.sweb.cz/) by [martinkunc](https://github.com/martinkunc/gobyexample-cz)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Italian](https://gobyexample.it) by the [Go Italian community](https://github.com/golangit/gobyexample-it)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Ukrainian](http://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)
* [Vietnamese](https://gobyexample.vn/) by [s6k Gopher](https://github.com/s6k-gopher/gobyexample-vn)

### Дякуємо

Дякуюємо [Jeremy Ashkenas](https://github.com/jashkenas)
за [Docco](http://jashkenas.github.com/docco/), що надихнули на цей проект.
