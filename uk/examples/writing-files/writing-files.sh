# Try running the file-writing code.
$ go run writing-files.go
Записано 5 байт
Записано 7 байт
Записано 9 байт

# Перевірте зміст записаних файлів.
$ cat /tmp/dat1
привіт
go
$ cat /tmp/dat2
some
writes
buffered

# Далі, ми розглянемо як застосовувати потоки
# вводу/виводу, на прикладі - створення програми
# "рядкового фільтра".
