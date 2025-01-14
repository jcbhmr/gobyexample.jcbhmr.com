# Запуск цієї програми призведе до її паніки, виводу
# повідомлення помилки та сліду горутини, та виходу з
# не нульовим статусом.

# When first panic in `main` fires, the program exits
# without reaching the rest of the code. If you'd like
# to see the program try to create a temp file, comment
# the first panic out.
$ go run panic.go
panic: маємо проблему

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# Зауважте що, на відміну від інших мов що
# використовують виключення (exceptions) для опрацювання
# помилок, в GO ідіоматично викорисовувати -
# _повернення помилок_, що вказуватимуть
# на проблеми як умога частіше.
