# Запустимо сервер у фоні.
$ go run context-in-http-servers.go &

# Симулюємо запита до `/hello`, невдовзі натиснемо
# Ctrl+C чим просигналізуємо про зупинку.
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
