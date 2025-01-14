# Запуск програми показує `FOO` матиме значення `1`
# (бо ми встановили його), `BAR` не матиме жодного
# значення, бо воно відсутнє в середовищі.
$ go run environment-variables.go
FOO: 1
BAR:

# Перелік ключів змінних середовища залежить від
# власне машини на якій код запускається.
TERM_PROGRAM
PATH
SHELL
...
FOO

# Якщо ми встановимо `BAR` в середовищі напередодні
# запуску, програма що запуститься слідом матиме
# змогу доступ до цієї зміннної.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
