# Скористайтесь цими командами щоб запустити
# приклад. (Зауваження: через обмеження
# тестового майданчику цей приклад може
# бути запущено лише на локальній машині.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

