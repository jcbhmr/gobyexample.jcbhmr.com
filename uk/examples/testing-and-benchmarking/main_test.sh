# Запуск усіх тестів у поточному пакеті у
# режимі "докладно".
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok  	examples/testing-and-benchmarking	0.023s

# Запуск усіх вимірювань у поточному проекті. Усі тести
# запустились до вимірювань, Прапорець `bench`
# відфільтровує імена функцій за допомогою регулярного
# виразу
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: examples/testing
BenchmarkIntMin-8 1000000000 0.3136 ns/op
PASS
ok  	examples/testing-and-benchmarking	0.351s
