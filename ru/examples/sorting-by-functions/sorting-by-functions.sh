# При запуске нашей программы отображается список,
# отсортированный по длине строки, как мы и хотели.
$ go run sorting-by-functions.go 
[kiwi peach banana]

# Следуя той же схеме создания пользовательского типа,
# реализации трех методов `интерфейса` для этого типа
# и последующего вызова `sort.Sort` для коллекции
# этого типа, мы можем сортировать срезы Go
# по произвольным функциям.
